package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"
	"time"

	"github.com/gogf/gf/v2/os/gcmd"
)

func pull_monitor() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	fid, err := client.EthNewBlockFilter(ctx)
	eth.Assert(err)
	fmt.Println("filter id: ", fid)

	index := 0
	ticker := time.Tick(2 * time.Second)
	for range ticker {
		hashes, err := client.EthGetFilterChanges(context.Background(), fid)
		eth.Assert(err)
		for _, hash := range hashes {
			fmt.Println("captured block hash: ", hash.Hex())
			blockInfo, err := client.EthGetBlockByHash(context.Background(), hash)
			eth.Assert(err)
			fmt.Println("blockInfo:", blockInfo)
		}

		if index == 10 {
			return
		}
		index++
	}
}

func ethToolsBlockMonitorPull() {
	fmt.Println("block monitor demo")
	go TriggerSendTransaction()
	pull_monitor()
}

var (
	EthToolsBlockMonitorPull = gcmd.Command{
		Name:  "ethToolsBlockMonitorPull",
		Usage: "ethToolsBlockMonitorPull",
		Brief: "eth tools block monitor pull  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsBlockMonitorPull()
			return
		},
	}
)
