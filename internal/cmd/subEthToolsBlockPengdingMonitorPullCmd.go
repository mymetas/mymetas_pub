package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"
	"time"

	"github.com/gogf/gf/v2/os/gcmd"
)

func pull_pending_monitor() {
	client, err := eth.Dial("ws://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	fid, err := client.EthNewPendingTransactionFilter(ctx)
	eth.Assert(err)

	index := 0
	timer := time.Tick(2 * time.Second)
	for range timer {
		hashes, err := client.EthGetFilterChanges(ctx, fid)
		eth.Assert(err)
		for _, hash := range hashes {
			fmt.Println("monitored pending txid: ", hash.Hex())

			//........error, becase this is pending block
			// blockInfo, err := client.EthGetBlockByHash(context.Background(), hash)
			// eth.Assert(err)
			// for _, tx := range blockInfo.Transactions() {
			// 	fmt.Println("tx:", tx)
			// 	fmt.Println("tx:", tx.Hash().Hex())
			// }
		}

		if index == 10 {
			return
		}

		index++
	}
}

func ethToolsBlockPendingMonitorPull() {
	fmt.Println("block monitor demo")
	go TriggerSendTransaction()
	pull_pending_monitor()
}

var (
	EthToolsBlockPendingMonitorPull = gcmd.Command{
		Name:  "ethToolsBlockPendingMonitorPull",
		Usage: "ethToolsBlockPendingMonitorPull",
		Brief: "eth tools block pending monitor pull  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsBlockPendingMonitorPull()
			return
		},
	}
)
