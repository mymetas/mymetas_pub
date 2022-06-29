package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/common"
)

// type RpcRequest struct {
// 	JsonRpc string        `json:"jsonrpc"`
// 	Method  string        `json:"method"`
// 	Params  []interface{} `json:"params"`
// 	Id      int64         `json:"id"`
// }

// type RpcResponse struct {
// 	JsonRpc string      `json:"jsonrpc"`
// 	Id      int64       `json:"id"`
// 	Result  interface{} `json:"result"`
// 	Error   interface{} `json:"error"`
// }

func DisplayBalance(ctx context.Context, client eth.Client, address []common.Address) {
	balance0, err := client.EthGetBalance(ctx, address[0], "latest")
	eth.Assert(err)
	fmt.Println("address[0]: ", address[0], " balance: ", balance0)
	balance1, err := client.EthGetBalance(ctx, address[1], "latest")
	eth.Assert(err)
	fmt.Println("address[1]: ", address[1], " balance: ", balance1)
}
