package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogf/gf/v2/os/gcmd"
)

func push_monitor() {
	client, err := eth.Dial("ws://localhost:8545")
	eth.Assert(err)

	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	eth.Assert(err)

	index := 0
	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case header := <-headers:
			fmt.Println("captured block hash: ", header.Hash().Hex())
			blockInfo, err := client.EthGetBlockByHash(context.Background(), header.Hash())
			eth.Assert(err)
			fmt.Println("block is:", blockInfo)
		}

		if index == 10 {
			return
		}
		index++
	}
}

func ethToolsBlockMonitorPush() {
	fmt.Println("block monitor demo")
	go TriggerSendTransaction()
	push_monitor()
}

var (
	EthToolsBlockMonitorPush = gcmd.Command{
		Name:  "ethToolsBlockMonitorPush",
		Usage: "ethToolsBlockMonitorPush",
		Brief: "eth tools block monitor Push  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsBlockMonitorPush()
			return
		},
	}
)
