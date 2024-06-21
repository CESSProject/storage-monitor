package core

import (
	"context"
	"github.com/CESSProject/cess-go-sdk/utils"
	"github.com/CESSProject/watchdog/constant"
	"github.com/CESSProject/watchdog/internal/model"
	"github.com/CESSProject/watchdog/internal/util"
	"log"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/pkg/errors"
)

func QueryMinerStatOnChain(accountId string, rpcAddr []string, mnemonic string, interval int) (model.MinerStat, error) {
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
	publicKey, err := utils.ParsingPublickey(accountId)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when parse public key")
	}
	chainInfo, err := chainClient.QueryMinerItems(publicKey, -1)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when query miner stat from chain")
	}
	stat, err = util.TransferMinerInfoToMinerStat(chainInfo)
	if err != nil {
		return model.MinerStat{}, errors.Wrap(err, "error when transfer data to stat object")
	}

	number, err := chainClient.QueryBlockNumber("")
	if err != nil {
		return stat, errors.Wrap(err, "error when query latest block number")
	}

	reward, err := chainClient.QueryRewardMap(publicKey, -1)
	if err != nil {
		log.Println("query reward failed: ", accountId)
		return stat, errors.Wrap(err, "error when query reward from chain")
	}

	stat.TotalReward = reward.TotalReward.Int
	stat.RewardIssued = reward.RewardIssued.Int

	if err != nil {
		return stat, err
	}
	latestBlockNum := int(number)
	blockUncheckNum := interval/constant.GenBlockInterval + 1 // 59/6 + 1 = 10
	for i := 0; i < blockUncheckNum; i++ {
		blockData, chainErr := chainClient.ParseBlockData(uint64(latestBlockNum - i))
		if chainErr != nil {
			log.Println("AccountId ", accountId, " query info from RPC: ", rpcAddr, " failed: ", chainErr)
		}
		punishmentInfo := blockData.Punishment
		for j := 0; j < len(punishmentInfo); j++ {
			stat.IsPunished[i][j] = punishmentInfo[j].From == accountId
		}
	}
	return stat, nil
}
