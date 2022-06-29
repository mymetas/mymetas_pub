package cmd

import (
	"context"
	"fmt"
	"math/big"
	"mymetas_pub/internal/service/eth"

	"github.com/gogf/gf/v2/os/gcmd"
)

func ethToolsGetBalance() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	addres, err := client.EthAccounts(ctx)
	eth.Assert(err)
	fmt.Println(addres)

	/*
		"latest"指最近的块数据，即当前的账户余额，也可指定特定的块
		"earliest"最初
		blockNumber := big.NewInt(12)
	*/
	balance1, err := client.EthGetBalance(ctx, addres[0], "latest")
	fmt.Println("latest", balance1)

	balance2, err := client.EthGetBalance(ctx, addres[0], "earliest")
	fmt.Println("earliest", balance2)

	blockByNumber0 := big.NewInt(0)
	balance3, err := client.EthGetBalance(ctx, addres[0], blockByNumber0)
	fmt.Println("blockNumber 0: ", balance3)

	blockByNumber1 := big.NewInt(1)
	balance4, err := client.EthGetBalance(ctx, addres[0], blockByNumber1)
	fmt.Println("blockNumber 1: ", balance4)
}

var (
	EthToolsGetBalance = gcmd.Command{
		Name:  "ethToolsGetBalance",
		Usage: "ethToolsGetBalance",
		Brief: "eth tools GetBalance  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsGetBalance()
			return
		},
	}
)
