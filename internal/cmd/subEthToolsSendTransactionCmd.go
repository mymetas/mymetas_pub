package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/gogf/gf/v2/os/gcmd"
)

func ethToolsSendTransaction() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	address, err := client.EthAccounts(ctx)
	eth.Assert(err)
	DisplayBalance(ctx, *client, address)

	fmt.Printf("transfer:  %v -> %v\n", address[0].Hex(), address[1].Hex())
	var msg = map[string]interface{}{
		"from":  address[0],
		"to":    address[1],
		"value": "0x100000",
		// "gas":      0x2000000,
		// "gasPrice": 0x0,
		"data": 0x1234,
		//"nonce": 5,
	}
	txid, err := client.EthSendTransaction(ctx, msg)
	eth.Assert(err)
	fmt.Println("txid: ", txid.Hex())

	// balance0, err := client.EthGetBalance(ctx, address[0], "latest")
	// fmt.Println("address[0]: ", address[0], " balance: ", balance0)
	// balance1, err := client.EthGetBalance(ctx, address[1], "latest")
	// fmt.Println("address[1]: ", address[1], " balance: ", balance1)
	DisplayBalance(ctx, *client, address)

	receipt, err := client.EthGetTransactionReceipt(ctx, txid)
	fmt.Println("receipt.transactionHash: ", receipt.TxHash)
	fmt.Println("receipt.transactionIndex: ", receipt.TransactionIndex)
	fmt.Println("receipt.blockHash: ", receipt.BlockHash)
	fmt.Println("receipt.blockNumber: ", receipt.BlockNumber)
	fmt.Println("receipt.CumulativeGasUsed: ", receipt.CumulativeGasUsed)
	fmt.Println("receipt.gasUsed: ", receipt.GasUsed)
	fmt.Println("receipt.contractAddrss: ", receipt.ContractAddress)
	fmt.Println("receipt.logs: ", receipt.Logs)
}

var (
	EthToolsSendTransaction = gcmd.Command{
		Name:  "ethToolsSendTransaction",
		Usage: "ethToolsSendTransaction",
		Brief: "ethTools SendTransaction  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsSendTransaction()
			return
		},
	}
)
