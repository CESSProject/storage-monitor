package core

import (
	"fmt"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"strconv"
	"sync"
	"time"

	"github.com/pkg/errors"
)

func (cli *WatchdogClient) SetChainData(signatureAcc string, rpcAddr []string, mnemonic string, interval int, miner string, created int64) (model.MinerStat, error) {
	var stat model.MinerStat
	hostIP := util.GetLocalIP()
	if hostIP == "" {
		hostIP = "127.0.0.1"
	}

	publicKey, err := utils.ParsingPublickey(signatureAcc)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when parse public key")
	}

	chainInfo, err := cli.CessChainClient.CessClient.QueryMinerItems(publicKey, -1)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when query miner stat from chain")
	}

	stat, err = util.TransferMinerInfoToMinerStat(chainInfo)
	if err != nil {
		log.Logger.Errorf("%s %s failed to transfer object format", hostIP, miner)
		return model.MinerStat{}, err
	}

	latestBlockNumberUint32, err := cli.CessChainClient.CessClient.QueryBlockNumber("")
	if err != nil {
		log.Logger.Errorf("%s %s failed to query latest block latestBlockNumberUint32", hostIP, miner)
		return stat, errors.Wrap(err, "failed to query latest block latestBlockNumberUint32")
	}

	if stat.Status != "positive" && time.Now().Unix()-created > 3600 {
		go alert(hostIP, miner, constant.MinerStatus, signatureAcc, "", latestBlockNumberUint32)
	}

	reward, err := cli.CessChainClient.CessClient.QueryRewardMap(publicKey, -1)
	if err != nil {
		log.Logger.Errorf("%s %s failed to query reward from chain", hostIP, miner)
		return stat, errors.Wrap(err, "failed to query reward from chain")
	}
	stat.TotalReward = util.BigNumConversion(reward.TotalReward)
	stat.RewardIssued = util.BigNumConversion(reward.RewardIssued)

	latestBlockNum := int(latestBlockNumberUint32)
	blockUncheckNum := interval / constant.GenBlockInterval
	latestUncheckBlockNum := latestBlockNum - blockUncheckNum

	scanApiUrl := fmt.Sprintf("%s/sminer/punishment?Acc=%s&pageindex=1&pagesize=1", constant.ScanApiUrl, signatureAcc)
	var response model.PunishSminerResponse

	err = cli.HTTPClient.Get(scanApiUrl, &response)
	if err != nil {
		log.Logger.Errorf("%s %s failed to query punishment from scan api", hostIP, miner)
		return stat, err
	}
	if response.Data.Count > 0 {
		scanApiUrl = fmt.Sprintf("%s/sminer/punishment?Acc=%s&pageindex=%d&pagesize=1", constant.ScanApiUrl, signatureAcc, response.Data.Count)
		err = cli.HTTPClient.Get(scanApiUrl, &response)
		if err != nil {
			log.Logger.Errorf("%s %s failed to query punishment from scan api", hostIP, miner)
			return stat, err
		}
		stat.LatestPunishInfo = response.Data.Content[0]
		if int(stat.LatestPunishInfo.BlockId) >= latestUncheckBlockNum {
			alertType := constant.SvcProofResIncorrect
			if stat.LatestPunishInfo.Type == 1 {
				alertType = constant.NoSubmitSvcProof
			}
			go alert(hostIP, miner, alertType, signatureAcc, stat.LatestPunishInfo.ExtrinsicHash, stat.LatestPunishInfo.BlockId)
		}
	}
	return stat, nil
}

func alert(hostIP string, miner string, alertType string, signatureAcc string, extrinsicHash string, blockId uint32) {
	var alertTypes = map[string]struct {
		description string
		url         string
	}{
		"Frozen":               {"The Storage Node status is Frozen", constant.ScanAccountURL},
		"NoSubmitSvcProof":     {"The Storage Node did not submit service file proof, you can check this block's system events in explorer", constant.ScanBlockURL},
		"SvcProofResIncorrect": {"The Storage Node service file proof checked by tee was incorrect", constant.ScanExtrinsicURL},
	}
	alertInfo, exists := alertTypes[alertType]
	if !exists {
		log.Logger.Warn("Unknown alert type")
		alertInfo = struct {
			description string
			url         string
		}{constant.DefaultDescription, constant.DefaultURL}
	}

	alertURL := alertInfo.url
	if alertType == "Frozen" {
		alertURL += signatureAcc
	} else if alertType == "NoSubmitSvcProof" {
		alertURL += strconv.Itoa(int(blockId))
	} else if alertType == "SvcProofResIncorrect" {
		alertURL += extrinsicHash
	}

	content := model.AlertContent{
		AlertTime:     time.Now().Format(constant.TimeFormat),
		HostIp:        hostIP,
		ContainerName: miner,
		Description:   alertInfo.description,
		DetailUrl:     alertURL,
	}
	log.Logger.Warnf("Triggered alert at block: %d, Alert content: %v", int(blockId), content)
	if !CustomConfig.Alert.Enable {
		return
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if SmtpConfig != nil {
			if err := SmtpConfig.SendMail(content); err != nil {
				log.Logger.Error("Failed to send alert email:", err)
			}
		}
	}()
	go func() {
		defer wg.Done()
		if WebhooksConfig != nil {
			if err := WebhooksConfig.SendAlertToWebhook(content); err != nil {
				log.Logger.Error("Failed to send alert webhook:", err)
			}
		}
	}()

	wg.Wait()
}
