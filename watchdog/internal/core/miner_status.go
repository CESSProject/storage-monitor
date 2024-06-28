package core

import (
	"context"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/log"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/pkg/errors"
)

func SetChainData(signatureAcc string, rpcAddr []string, mnemonic string, interval int, host string, miner string) (model.MinerStat, error) {
	var stat model.MinerStat
	chainClient, err := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(rpcAddr),
		cess.TransactionTimeout(time.Second*30),
		cess.Mnemonic(mnemonic),
	)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when new a cess client")
	}
	publicKey, err := utils.ParsingPublickey(signatureAcc)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when parse public key")
	}
	chainInfo, err := chainClient.QueryMinerItems(publicKey, -1)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when query miner stat from chain")
	}
	stat, err = util.TransferMinerInfoToMinerStat(chainInfo)
	if stat.Status != "positive" {
		alert(stat, host, miner)
	}
	if err != nil {
		log.Logger.Errorf("%s %s failed to transfer object format", host, miner)
		return model.MinerStat{}, err
	}

	number, err := chainClient.QueryBlockNumber("")
	if err != nil {
		log.Logger.Errorf("%s %s failed to query latest block number", host, miner)
		return stat, err
	}

	reward, err := chainClient.QueryRewardMap(publicKey, -1)
	if err != nil {
		log.Logger.Errorf("%s %s failed to query reward from chain", host, miner)
		return stat, err
	}

	stat.TotalReward = util.BigNumConversion(reward.TotalReward)
	stat.RewardIssued = util.BigNumConversion(reward.RewardIssued)

	latestBlockNum := int(number)
	blockUncheckNum := interval/constant.GenBlockInterval + 1 // 59/6 + 1 = 10
	for i := 0; i < blockUncheckNum; i++ {
		blockData, chainErr := chainClient.ParseBlockData(uint64(latestBlockNum - i))
		if chainErr != nil {
			log.Logger.Errorf("%s %s failed to query info from rpc %s", host, miner, rpcAddr)
		}
		punishmentInfo := blockData.Punishment
		for j := 0; j < len(punishmentInfo); j++ {
			stat.IsPunished[i][j] = punishmentInfo[j].From == signatureAcc
			if stat.IsPunished[i][j] {
				alert(stat, host, miner)
			}
		}
	}
	return stat, nil
}

func alert(stat model.MinerStat, host string, miner string) {
	if CustomConfig.Alert.Enable {
		content := model.AlertContent{
			AlertTime:     time.Now().Format(constant.TimeFormat),
			HostIp:        host,
			ContainerName: miner,
			Description:   "The Storage Miner is not a positive status or get punishment",
		}
		go func() {
			if SmtpConfigPoint != nil {
				err := SmtpConfigPoint.SendMail(content)
				if err != nil {
					return
				}
			}
		}()
		go func() {
			if WebhookConfigPoint != nil {
				err := WebhookConfigPoint.SendAlertToWebhook(content)
				if err != nil {
					return
				}
			}
		}()
		log.Logger.Errorf("%s %s status is not a positive status: %s", host, miner, stat.Status)
	}
}
