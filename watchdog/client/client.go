package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/CESSProject/watchdog/status"
	"github.com/btcsuite/btcutil/base58"
	"github.com/centrifuge/go-substrate-rpc-client/v4/signature"
	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
)

const (
	DEFAULT_PUSH_FREQ = time.Second * 30
	DEFAULT_PUSH_PATH = "/push"
)

type MinerInfo struct {
	AccountId string
	Conf      status.Configfile
	CInfo     status.Container
	CStat     chan status.ContainerStat
	ChainInfo pattern.MinerInfo
}

type MinerInfoDisplay struct {
	ContainerInfo status.Container     `json:"container_info"`
	ContainerStat status.ContainerStat `json:"container_status"`
	Metadata      status.MinerMetadata `json:"miner_metadata"`
}

type WatchdogClient struct {
	*status.Client
	clist map[string]*MinerInfo
}

var watchdog *WatchdogClient
var exeConf = types.ExecConfig{
	Cmd:          []string{"cat", "config.yaml"},
	WorkingDir:   "/opt/bucket/",
	AttachStdout: true,
	AttachStderr: true,
	Tty:          true,
}

type Watchdog interface {
}

func GetWatchdogClient() Watchdog {
	return watchdog
}

func InitWatchdogClient() error {

	if watchdog == nil {
		watchdog = &WatchdogClient{}
	}
	// create docker client
	cli, err := status.NewClient()
	if err != nil {
		return errors.Wrap(err, "init watchdog client error")
	}
	watchdog.Client = cli

	//
	list, err := cli.ListContainers()
	if err != nil {
		return errors.Wrap(err, "init watchdog client error")
	}
	watchdog.clist = make(map[string]*MinerInfo)

	//get container info and miner config
	for _, c := range list {
		log.Println("check container", c.ID[:8], "image", c.Image)
		if !strings.Contains(c.Image, status.MINER_IMAGE) {
			continue
		}
		err := watchdog.InsertMinerInfo(c)
		if err != nil {
			return errors.Wrap(err, "init watchdog client error")
		}
	}

	//get miners' status
	for _, v := range watchdog.clist {
		err := cli.ContainerStats(context.Background(), v.CInfo.ID, v.CStat)
		if err != nil {
			return errors.Wrap(err, "init watchdog client error")
		}
	}

	// get miners' info on chain
	for _, v := range watchdog.clist {
		info, err := status.QueryMinerInfoOnChain(v.AccountId, v.Conf.Rpc)
		if err != nil {
			return errors.Wrap(err, "init watchdog client error")
		}
		v.ChainInfo = info
	}

	return nil
}

func (cli *WatchdogClient) DeleteMinerInfo(cid string) {
	if value, ok := cli.clist[cid]; ok {
		close(value.CStat)
		delete(cli.clist, cid)
	}
}

func (cli *WatchdogClient) InsertMinerInfo(cinfo status.Container) error {
	res, err := cli.ExeCommand(cinfo.ID, exeConf)
	if err != nil {
		return errors.Wrap(err, "insert miner info error")
	}

	conf, err := status.ParseMinerConfigFile(res[8:])
	if err != nil {
		return errors.Wrap(err, "insert miner info error")
	}

	key, err := signature.KeyringPairFromSecret(conf.Mnemonic, 0)
	if err != nil {
		return errors.Wrap(err, "insert miner info error")
	}

	acc, err := utils.EncodePublicKeyAsCessAccount(key.PublicKey)
	if err != nil {
		return errors.Wrap(err, "insert miner info error")
	}

	info := MinerInfo{
		AccountId: acc,
		CInfo:     cinfo,
		Conf:      conf,
		CStat:     make(chan status.ContainerStat, 1),
	}

	cli.clist[cinfo.ID] = &info
	return nil
}

func (cli *WatchdogClient) FlashContainerList() error {
	list, err := cli.ListContainers()
	if err != nil {
		return errors.Wrap(err, "init watchdog client error")
	}

	mp := make(map[string]struct{})
	for _, c := range list {
		if !strings.Contains(c.Image, status.MINER_IMAGE) {
			continue
		}
		mp[c.ID] = struct{}{}
		if v, ok := cli.clist[c.ID]; ok {
			v.CInfo = c
			continue
		}
		err := cli.InsertMinerInfo(c)
		if err != nil {
			return errors.Wrap(err, "flash container list error")
		}
		err = cli.ContainerStats(context.Background(), c.ID, cli.clist[c.ID].CStat)
		if err != nil {
			return errors.Wrap(err, "init watchdog client error")
		}
	}

	for _, v := range cli.clist {

		//clean
		if _, ok := mp[v.CInfo.ID]; !ok {
			//cli.DeleteMinerInfo(v.CInfo.ID)
			continue
		}

		//updata container status
		<-v.CStat

		// update miner info on chain
		info, err := status.QueryMinerInfoOnChain(v.AccountId, v.Conf.Rpc)
		if err != nil {
			return errors.Wrap(err, "init watchdog client error")
		}
		v.ChainInfo = info
	}

	return nil
}

func (cli *WatchdogClient) ExportMinerInfo() []MinerInfoDisplay {
	dataList := make([]MinerInfoDisplay, 0, len(cli.clist))
	for _, v := range cli.clist {
		data := MinerInfoDisplay{}
		data.ContainerInfo = v.CInfo
		data.ContainerStat = <-v.CStat
		data.Metadata = status.MinerMetadata{
			Name:            "storage miner",
			PeerId:          base58.Encode([]byte(string(v.ChainInfo.PeerId[:]))),
			State:           string(v.ChainInfo.State),
			StakingAmount:   v.ChainInfo.Collaterals.String(),
			ValidatedSpace:  v.ChainInfo.IdleSpace.Uint64(),
			UsedSpace:       v.ChainInfo.ServiceSpace.Uint64(),
			LockedSpace:     v.ChainInfo.LockSpace.Uint64(),
			StakingAccount:  v.AccountId,
			EarningsAccount: v.Conf.EarningsAcc,
		}
		dataList = append(dataList, data)
	}
	return dataList
}

func RunWatchdogClient(url string, freq time.Duration) error {
	if watchdog == nil {
		err := errors.New("watchdog client is not initialized")
		return errors.Wrap(err, "run watchdog client error")
	}
	if freq <= 0 {
		freq = DEFAULT_PUSH_FREQ
	}

	for {
		data := watchdog.ExportMinerInfo()
		err := PostDataToServer(url+DEFAULT_PUSH_PATH, data)
		if err != nil {
			log.Println("run watchdog client error:", err)
		}
		time.Sleep(freq)
		err = watchdog.FlashContainerList()
		if err != nil {
			log.Println("run watchdog client error:", err)
		}
	}
}

func PostDataToServer(url string, data []MinerInfoDisplay) error {

	jbytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "post data to server error")
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jbytes))
	if err != nil {
		return errors.Wrap(err, "post data to server error")
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return errors.Wrap(err, "post data to server error")
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "post data to server error")
	}
	log.Println("watchdog server response:", respData)
	return nil
}
