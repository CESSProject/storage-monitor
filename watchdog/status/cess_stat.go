package status

import (
	"context"
	"time"

	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/core/pattern"
	"github.com/CESSProject/cess-go-sdk/core/utils"
	"github.com/pkg/errors"
)

type MinerMetadata struct {
	Name            string `json:"name"`
	PeerId          string `json:"peer_id"`
	State           string `json:"state"`
	StakingAmount   string `json:"staking_amount"`
	ValidatedSpace  uint64 `json:"validated_space"`
	UsedSpace       uint64 `json:"used_space"`
	LockedSpace     uint64 `json:"locked_space"`
	StakingAccount  string `json:"staking_account"`
	EarningsAccount string `json:"earnings_account"`
}

func QueryMinerInfoOnChain(accountId string, rpcaddrs []string) (pattern.MinerInfo, error) {
	var info pattern.MinerInfo
	sdk, err := cess.New(
		context.Background(),
		"bucket-watchdog",
		cess.ConnectRpcAddrs(rpcaddrs),
		cess.TransactionTimeout(time.Second*30),
	)
	if err != nil {
		return info, errors.Wrap(err, "query miner info on chain error")
	}

	pubkey, err := utils.ParsingPublickey(accountId)
	if err != nil {
		return info, errors.Wrap(err, "query miner info on chain error")
	}
	info, err = sdk.QueryStorageMiner(pubkey)
	if err != nil {
		return info, errors.Wrap(err, "query miner info on chain error")
	}
	return info, nil
}
