package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/os/gcmd"
)

func push_pending_monitor() {
	client, err := eth.Dial("ws://localhost:8545")
	eth.Assert(err)

	hashes := make(chan *common.Hash)
	sub, err := client.SubscribeNewPendingTransactions(context.Background(), hashes)
	eth.Assert(err)

	index := 0
	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case hash := <-hashes:
			fmt.Println("captured tx id: ", hash.Hex())
		}

		if index == 10 {
			return
		}
		index++
	}
}

func ethToolsBlockPendingMonitorPush() {
	fmt.Println("block monitor demo")
	go TriggerSendTransaction()
	push_pending_monitor()
}

var (
	EthToolsBlockPendingMonitorPush = gcmd.Command{
		Name:  "ethToolsBlockPendingMonitorPush",
		Usage: "ethToolsBlockPendingMonitorPush",
		Brief: "eth tools block pending monitor Push  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsBlockPendingMonitorPush()
			return
		},
	}
)
