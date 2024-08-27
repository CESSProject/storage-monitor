package core

import (
	"context"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/docker/docker/api/types"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var exeConf = types.ExecConfig{
	Cmd:          []string{"cat", "config.yaml"},
	WorkingDir:   "/opt/miner/",
	AttachStdout: true,
	AttachStderr: true,
}

type WatchdogClient struct {
	Host                  string                // 127.0.0.1 or some ip else
	*Client                                     // docker cli
	*util.HTTPClient                            // http cli
	*util.CessChainClient                       // cess chain cli
	MinerInfoMap          map[string]*MinerInfo // key: miner-name
	Updating              bool                  // is miners data updating ?
	Active                bool                  // sleep or run
	mutex                 sync.Mutex
}

var Clients = map[string]*WatchdogClient{} // key: hostIP

type MinerInfo struct {
	Name         string
	SignatureAcc string
	Conf         model.MinerConfigFile
	CInfo        model.Container
	MinerStat    model.MinerStat
}

func InitWatchdogClients(conf model.YamlConfig) error {
	hosts := conf.Hosts
	Clients = make(map[string]*WatchdogClient, len(hosts))
	var initClientsWG sync.WaitGroup
	errChan := make(chan error, len(hosts))
	for _, host := range hosts {
		initClientsWG.Add(1)
		go func(host model.HostItem) {
			defer initClientsWG.Done()
			dockerClient, err := NewClient(host)
			httpClient := util.NewHTTPClient()
			chainClient := util.NewCessChainClient([]string{constant.DefaultRpcUrl, constant.LocalRpcUrl})
			if dockerClient == nil {
				return
			}
			if err != nil {
				errChan <- err
				return
			}
			Clients[host.IP] = &WatchdogClient{
				Host:            host.IP,
				Client:          dockerClient,
				HTTPClient:      httpClient,
				CessChainClient: chainClient,
				MinerInfoMap:    make(map[string]*MinerInfo),
				Updating:        false,
				Active:          true,
			}
			log.Logger.Infof("Create a docker client with host: %s successfully", host.IP)
		}(host)
	}
	initClientsWG.Wait()
	close(errChan)
	for err := range errChan {
		if err != nil {
			return err
		}
	}
	log.Logger.Info("Init All Watchdog Clients Successfully")
	return nil
}

func RunWatchdogClients(conf model.YamlConfig) error {
	for hostIp, client := range Clients {
		if client == nil {
			log.Logger.Warnf("Client for host %s is nil, skipping", hostIp)
			continue
		}
		log.Logger.Infof("Start to run task at host: %s", hostIp)
		go client.RunWatchdogClient(conf)
	}
	return nil
}

func (cli *WatchdogClient) RunWatchdogClient(conf model.YamlConfig) {
	for cli.Active {
		if err := cli.start(conf); err != nil {
			log.Logger.Warnf("Error when start %s watchdog client %v", cli.Host, err)
		}
		time.Sleep(time.Duration(CustomConfig.ScrapeInterval) * time.Second) // scrape interval
	}
	// gc will collect inactive clients
}

func (cli *WatchdogClient) start(conf model.YamlConfig) error {
	cli.Updating = true
	defer func() { cli.Updating = false }()
	interval := conf.ScrapeInterval
	ctx := context.Background()
	containers, err := cli.Client.ListContainers(ctx)

	if err != nil {
		log.Logger.Errorf("Error when listing %s containers: %v", cli.Host, err)
		return err
	}

	errChan := make(chan error, len(containers))
	done := make(chan struct{})
	go func() {
		for err := range errChan {
			log.Logger.Errorf("Error when %s task run: %v", cli.Host, err)
		}
		close(done)
	}()

	// get miner info and miner config
	var setContainersDataWG sync.WaitGroup
	for _, container := range containers {
		if !strings.Contains(container.Image, constant.MinerImage) {
			continue
		}
		runningMiners := make(map[string]bool)
		for _, v := range containers {
			runningMiners[v.Name] = true
		}
		for key := range cli.MinerInfoMap {
			// delete un-exist miner
			if !runningMiners[key] {
				delete(cli.MinerInfoMap, key)
			}
		}
		setContainersDataWG.Add(1)
		go func(container model.Container) {
			defer setContainersDataWG.Done()
			if err := cli.SetContainerData(ctx, container); err != nil {
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
			if res, err := cli.SetContainerStats(ctx, m.CInfo.ID); err != nil {
				errChan <- err
			} else {
				m.CInfo.CPUPercent = res.CPUPercent
				m.CInfo.MemoryPercent = res.MemoryPercent
				m.CInfo.MemoryUsage = res.MemoryUsage
			}
		}(miner)
	}
	setContainersStatsDataWG.Wait()

	// set miner's info on chain
	// qps might be high when use goroutine to request scan server,
	for _, miner := range cli.MinerInfoMap {
		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		sleepDuration := r.Intn(2) + 1
		time.Sleep(time.Duration(sleepDuration) * time.Second)
		if minerStat, err := cli.SetChainData(miner.SignatureAcc, miner.Conf.Rpc, miner.Conf.Mnemonic, interval, miner.Name, miner.CInfo.Created); err != nil {
			errChan <- err
		} else {
			miner.MinerStat = minerStat
			cli.MinerInfoMap[miner.Name].MinerStat = miner.MinerStat
		}
	}

	close(errChan)
	<-done
	return nil
}

func (cli *WatchdogClient) SetContainerData(ctx context.Context, cinfo model.Container) error {
	res, err := cli.ExeCommand(ctx, cinfo.ID, exeConf)
	if err != nil {
		log.Logger.Errorf("%s read config from container path: %s failed in host: %s", cinfo.Name, constant.MinerConfPath, cli.Host)
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
		log.Logger.Errorf("Failed to parse storage node config file for container %s: %v", cinfo.Name, err)
		return err
	}

	key, err := signature.KeyringPairFromSecret(conf.Mnemonic, 0)
	if err != nil {
		log.Logger.Errorf("Failed to generate keyring pair for container %s: %v", cinfo.Name, err)
		return err
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(key.PublicKey)
	if err != nil {
		log.Logger.Errorf("Failed to encode public key as Cess account for container %s: %v", cinfo.Name, err)
		return err
	}
	cli.mutex.Lock()
	defer cli.mutex.Unlock()
	if miner, ok := cli.MinerInfoMap[cinfo.Name]; ok {
		miner.Name = conf.Name
		miner.SignatureAcc = acc
		miner.Conf = conf
		miner.CInfo = cinfo
	} else {
		cli.MinerInfoMap[cinfo.Name] = &MinerInfo{
			Name:         conf.Name,
			SignatureAcc: acc,
			CInfo:        cinfo,
			Conf:         conf,
			MinerStat:    model.MinerStat{},
		}
	}
	return nil
}
