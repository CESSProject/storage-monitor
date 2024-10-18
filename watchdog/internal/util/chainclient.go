package util

import (
	"context"
	cess "github.com/CESSProject/cess-go-sdk"
	"github.com/CESSProject/cess-go-sdk/chain"
	"time"
)

type CessChainClient struct {
	CessClient *chain.ChainClient
}

func NewCessChainClient(rpcUrl []string) *CessChainClient {
	chainClient, _ := cess.New(
		context.Background(),
		cess.ConnectRpcAddrs(rpcUrl),
		cess.TransactionTimeout(time.Second*15),
	)
	return &CessChainClient{CessClient: chainClient}
}
