package core

import (
	"context"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"strings"
	"sync"
	"time"

	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/docker/docker/api/types"
)

var exeConf = types.ExecConfig{
	Cmd:          []string{"cat", "config.yaml"},
	WorkingDir:   "/opt/miner/",
	AttachStdout: true,
	AttachStderr: true,
}

type WatchdogClient struct {
	Host         string                // 127.0.0.1
	*Client                            // docker cli
	MinerInfoMap map[string]*MinerInfo // key: miner-name
}

var Clients = map[string]*WatchdogClient{} // key: hostIP

type MinerInfo struct {
	Name      string
	AccountId string
	Conf      model.MinerConfigFile
	CInfo     model.Container
	MinerStat model.MinerStat
}

func InitWatchdogClients(conf model.YamlConfig) error {
	hosts := conf.Hosts
	Clients = make(map[string]*WatchdogClient, len(hosts))
	var initClientsWG sync.WaitGroup
	errChan := make(chan error, len(hosts))
	for idx, host := range hosts {
		initClientsWG.Add(1)
		go func(idx int, host model.HostItem) {
			defer initClientsWG.Done()
			dockerClient, err := NewClient(host)
			if dockerClient == nil {
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			log.Logger.Infof("Create a docker client with host: %s successfully", host.IP)
			Clients[host.IP] = &WatchdogClient{
				Host:         host.IP,
				Client:       dockerClient,
				MinerInfoMap: make(map[string]*MinerInfo),
			}
		}(idx, host)
	}
	initClientsWG.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	return nil
}

func RunWatchdogClients(conf model.YamlConfig) error {
	for hostIp, client := range Clients {
		if client != nil {
			log.Logger.Infof("Start to run task at %s", hostIp)
			go client.RunWatchdogClient(conf)
		}
	}
	return nil
}

func (cli *WatchdogClient) RunWatchdogClient(conf model.YamlConfig) {
	interval := conf.ScrapeInterval
	for {
		err := cli.start(conf)
		if err != nil {
			log.Logger.Fatalf("Error when start watchdog %v", err)
		}
		time.Sleep(time.Duration(interval) * time.Second) // scrape interval
	}
}

func (cli *WatchdogClient) start(conf model.YamlConfig) error {
	interval := conf.ScrapeInterval
	containers, err := cli.Client.ListContainers()

	if err != nil {
		log.Logger.Errorf("Error when listing %s containers: %v", cli.Host, err)
		return err
	}

	errChan := make(chan error, len(containers))
	done := make(chan bool)
	go func() {
		for err = range errChan {
			log.Logger.Errorf("Error when %s task run: %v", cli.Host, err)
		}
		done <- true
	}()

	// get miner info and miner config
	var setContainersDataWG sync.WaitGroup
	for _, container := range containers {
		if !strings.Contains(container.Image, constant.MinerImage) {
			continue
		}
		log.Logger.Infof("Task Start: %s, Miner: %s", cli.Host, container.Name)
		setContainersDataWG.Add(1)
		go func(container model.Container) {
			defer setContainersDataWG.Done()
			err = cli.SetContainerData(container)
			if err != nil {
				errChan <- err
			}
		}(container)
	}
	setContainersDataWG.Wait()

	// set miners' container stats
	var setContainersStatsDataWG sync.WaitGroup
	for _, miner := range cli.MinerInfoMap {
		setContainersStatsDataWG.Add(1)
		go func(m *MinerInfo) {
			defer setContainersStatsDataWG.Done()
			res, err := cli.SetContainerStats(context.Background(), m.CInfo.ID)
			if err != nil {
				errChan <- err
				return
			}
			m.CInfo.CPUPercent = res.CPUPercent
			m.CInfo.MemoryPercent = res.MemoryPercent
			m.CInfo.MemoryUsage = res.MemoryUsage
		}(miner)
	}
	setContainersStatsDataWG.Wait()

	// set miner's info on chain
	var setChainDataWG sync.WaitGroup
	for _, miner := range cli.MinerInfoMap {
		setChainDataWG.Add(1)
		go func(m *MinerInfo) {
			defer setChainDataWG.Done()
			m.MinerStat, _ = SetChainData(m.AccountId, m.Conf.Rpc, m.Conf.Mnemonic, interval, cli.Host, m.Name)
			cli.MinerInfoMap[m.Name].MinerStat = m.MinerStat
			log.Logger.Infof("Task Done: %s, Miner: %s", cli.Host, m.Name)
		}(miner)
	}
	setChainDataWG.Wait()
	close(errChan)
	<-done
	return nil
}

func (cli *WatchdogClient) SetContainerData(cinfo model.Container) error {
	res, err := cli.ExeCommand(cinfo.ID, exeConf)
	if err != nil {
		log.Logger.Errorf("%s read config from /opt/miner/config.yaml failed in host: %s", cinfo.Name, cli.Host)
		return err
	}

	// res:
	//[1 0 0 0 0 0 1 179 78 97 109 101 58 32 109 105
	//110 101 114 49 13 10 80 111 114 116 58 32 49 53
	//48 48 49 13 10 69 97 114 110 105 110 103 115 65
	//114 46 98 111 111 116 45 109 105 110 101 114 45
	//100 101 118 110 101 116 46 99 101 115 115 46 99]

	// 1-8 byte: 1 0 0 0 0 0 1 179

	// The First Byte：stream dataType. value: 0x01 stdout，value: 0x02 stderr

	// The second to fourth bytes are 0 0 0: Reserved and not used

	// The fifth to eighth bytes are 0 0 1 179: they indicate the length of the following data block.
	//The length here is a 32-bit integer with a value of 0x000001B3, which is 435 in decimal.
	conf, err := util.ParseMinerConfigFile(res[8:]) // delete 0-7
	if err != nil {
		return err
	}

	key, err := signature.KeyringPairFromSecret(conf.Mnemonic, 0)
	if err != nil {
		return err
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(key.PublicKey)
	if err != nil {
		return err
	}

	info := MinerInfo{
		Name:      conf.Name,
		AccountId: acc,
		CInfo:     cinfo,
		Conf:      conf,
		MinerStat: model.MinerStat{},
	}
	cli.MinerInfoMap[cinfo.Name] = &info
	return nil
}
