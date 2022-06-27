package cmd

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gogf/gf/v2/os/gcmd"
)

func blockByNumber() {
	client, err := ethclient.Dial("http://localhost:8545")
	assert(err)
	blockNumber := big.NewInt(0)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	assert(err)
	fmt.Println("hash: ", block.Hash().Hex())
	fmt.Println("coinbase: ", block.Coinbase().Hex())
	fmt.Println("num of transactions: ", block.Transactions().Len())
}

func gasPrice() {
	client, err := ethclient.Dial("http://localhost:8545")
	assert(err)
	price, err := client.SuggestGasPrice(context.Background())
	assert(err)
	fmt.Println("gas price: ", price)
}

var (
	EthClient = gcmd.Command{
		Name:  "ethClient",
		Usage: "ethClient",
		Brief: "eth client  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			blockByNumber()
			gasPrice()
			return
		},
	}
)
