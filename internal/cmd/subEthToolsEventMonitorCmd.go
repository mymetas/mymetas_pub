package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"mymetas_pub/internal/service/eth"
	"mymetas_pub/internal/service/eth/contract/wrapper/testerc20token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/os/gcmd"
)

func event_monitor() {
	addrHexBytes, err := ioutil.ReadFile(HappyTokenAddrPath)
	eth.Assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	eth.Assert(err)

	client, err := eth.Dial("ws://localhost:8545")
	eth.Assert(err)

	inst, err := testerc20token.NewTesterc20token(contractAddress, client)
	eth.Assert(err)

	watchOpts := &bind.WatchOpts{}
	events := make(chan *testerc20token.Testerc20tokenTransfer)
	_from := []common.Address{}
	_to := []common.Address{}
	sub, err := inst.WatchTransfer(watchOpts, events, _from, _to)
	eth.Assert(err)
	//fmt.Println(sub)

	index := 0
	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case event := <-events:
			fmt.Println("captured:")
			fmt.Println("-> from: ", event.From.Hex())
			fmt.Println("-> to: ", event.To.Hex())
			fmt.Println("-> value:", event.Value)
		}

		if index >= 10 {
			return
		}
		index++
	}
}

func ethToolsEventMonitor() {
	fmt.Println("events monitor  demo")
	go TriggerSendContractTransaction()
	event_monitor()
}

var (
	EthToolsEventMonitor = gcmd.Command{
		Name:  "ethToolsEventMonitor",
		Usage: "ethToolsEventMonitor",
		Brief: "eth tools event  monitor   ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsEventMonitor()
			return
		},
	}
)
