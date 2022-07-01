package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/gogf/gf/v2/os/gcmd"
)

func push_logs_monitor1() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	opts := map[string]string{}
	fid, err := client.EthNewFilter(ctx, opts)
	eth.Assert(err)
	fmt.Println("filter id: ", fid)

	index := 0
	ticker := time.Tick(2 * time.Second)
	for range ticker {
		logs, err := client.EthGetLogFilterChanges(ctx, fid)
		eth.Assert(err)
		for _, log := range logs {
			fmt.Printf("captured log:%+v\n", log)
		}

		if index >= 10 {
			return
		}
		index++
	}
}

func push_logs_monitor() {
	client, err := eth.Dial("ws://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	query := &ethereum.FilterQuery{}
	logs := make(chan types.Log)
	sub, _ := client.SubscribeFilterLogs(
		ctx,
		*query,
		logs)

	index := 0
	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case log := <-logs:
			fmt.Println("log txid:", log.TxHash.Hex())
		}

		if index >= 10 {
			return
		}
		index++
	}
}

func ethToolsLogsMonitorPush() {
	fmt.Println("logs monitor push demo")
	go TriggerSendContractTransaction()
	push_logs_monitor()
}

var (
	EthToolsLogsMonitorPush = gcmd.Command{
		Name:  "ethToolsLogsMonitorPush",
		Usage: "ethToolsLogsMonitorPush",
		Brief: "eth tools logs  monitor Push  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsLogsMonitorPush()
			return
		},
	}
)
