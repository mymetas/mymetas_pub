package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"mymetas_pub/internal/service/eth"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var AbiBytesPath = "internal/service/eth/contract/build/TestErc20Token.abi"
var BinHexBytesPath = "internal/service/eth/contract/build/TestErc20Token.bin"
var HappyTokenAddrPath = "internal/service/eth/contract/HappyToken.addr"

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

func TriggerSendTransaction() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	eth.Assert(err)

	timer := time.Tick(5 * time.Second)
	for range timer {
		msg := map[string]interface{}{
			"from":  accounts[0],
			"to":    accounts[1],
			"value": big.NewInt(1000),
		}
		txid, err := client.EthSendTransaction(context.Background(), msg)
		eth.Assert(err)
		fmt.Println("trigger txid: ", txid.Hex())
	}
}

func TriggerSendContractTransaction() {
	abiBytes, err := ioutil.ReadFile(AbiBytesPath)
	eth.Assert(err)
	tokenAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	eth.Assert(err)

	addrBytes, err := ioutil.ReadFile(HappyTokenAddrPath)
	eth.Assert(err)
	contractAddress := common.HexToAddress(string(addrBytes))

	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	eth.Assert(err)

	data, err := tokenAbi.Pack("transfer", accounts[1], big.NewInt(100))
	eth.Assert(err)
	msg := map[string]interface{}{
		"from": accounts[0],
		"to":   contractAddress,
		//"data": common.ToHex(data),
		"data": "0x" + common.Bytes2Hex(data),
		"gas":  big.NewInt(4000000),
	}

	ticker := time.Tick(5 * time.Second)
	for range ticker {
		txid, err := client.EthSendTransaction(ctx, msg)
		eth.Assert(err)
		fmt.Println("trigger txid: ", txid.Hex())
	}
}
