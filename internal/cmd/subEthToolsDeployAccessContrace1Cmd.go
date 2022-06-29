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
	"github.com/gogf/gf/v2/os/gcmd"
)

var abiBytesPath = "internal/service/eth/contract/build/TestErc20Token.abi"
var binHexBytesPath = "internal/service/eth/contract/build/TestErc20Token.bin"
var happyTokenAddrPath = "internal/service/eth/contract/HappyToken.addr"

func deployHappyTokenContract() {
	fmt.Println("deploy contract theory demo")
	abiBytes, err := ioutil.ReadFile(abiBytesPath)
	eth.Assert(err)
	//fmt.Println(abiBytes)

	binHexBytes, err := ioutil.ReadFile(binHexBytesPath)
	eth.Assert(err)
	binBytes := common.Hex2Bytes(string(binHexBytes))
	//fmt.Println(binBytes)

	tokenAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	eth.Assert(err)
	encodedParams, err := tokenAbi.Pack(
		"",
		big.NewInt(1000000),
		"HAPPY COIN",
		uint8(0),
		"HAPY",
	)
	eth.Assert(err)
	//fmt.Println(encodedParams)
	data := append(binBytes, encodedParams...)

	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	eth.Assert(err)

	msg := map[string]interface{}{
		"from": accounts[0],
		//"data": common.ToHex(data),
		"data": "0x" + common.Bytes2Hex(data),
		"gas":  big.NewInt(2000000),
	}
	txid, err := client.EthSendTransaction(ctx, msg)
	eth.Assert(err)
	fmt.Println("txid: ", txid.Hex())

	//wait
	time.Sleep(time.Second)

	receipt, err := client.EthGetTransactionReceipt(ctx, txid)
	eth.Assert(err)
	fmt.Println("contract address: ", receipt.ContractAddress.Hex())
	err = ioutil.WriteFile(happyTokenAddrPath, []byte(receipt.ContractAddress.Hex()), 0644)
	eth.Assert(err)
	fmt.Println("done.")
}

func accessContractTransfer() {
	fmt.Println("access contract theory demo")

	abiBytes, err := ioutil.ReadFile(abiBytesPath)
	eth.Assert(err)
	tokenAbi, err := abi.JSON(strings.NewReader(string(abiBytes)))
	eth.Assert(err)

	addrBytes, err := ioutil.ReadFile(happyTokenAddrPath)
	eth.Assert(err)
	contractAddress := common.HexToAddress(string(addrBytes))
	fmt.Println("contract address: ", contractAddress.Hex())

	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	accounts, err := client.EthAccounts(ctx)
	eth.Assert(err)

	//transfer from account 0 => 1

	data, err := tokenAbi.Pack(
		"transfer",
		accounts[1],
		big.NewInt(10000),
	)
	eth.Assert(err)
	msg := map[string]interface{}{
		"from": accounts[0],
		"to":   contractAddress,
		"gas":  big.NewInt(2000000),
		//"data": common.ToHex(data),
		"data": "0x" + common.Bytes2Hex(data),
	}

	txid, err := client.EthSendTransaction(ctx, msg)
	fmt.Println("txid is ", txid)

	encodeParams, err := tokenAbi.Pack(
		//"balanceof", //error, function name fail
		"balanceOf", //ok
		accounts[0],
	)

	msg1 := map[string]interface{}{
		"from": accounts[0],
		"to":   contractAddress,
		"data": "0x" + common.Bytes2Hex(encodeParams),
		"gas":  big.NewInt(2000000),
	}

	balance, err := client.EthCall(ctx, msg1)
	eth.Assert(err)
	fmt.Println("balance: ", balance)

	balanceOut, err1 := tokenAbi.Unpack("balanceOf", balance)
	eth.Assert(err1)
	fmt.Println("balance decoded: ", balanceOut)
}

func ethToolsDeployAccessContract1() {
	deployHappyTokenContract()
	accessContractTransfer()
}

var (
	EthToolsDeployAccessContract1 = gcmd.Command{
		Name:  "ethToolsDeployAccessContract1",
		Usage: "ethToolsDeployAccessContract1",
		Brief: "ethTools Deploy && AccessContract 1  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsDeployAccessContract1()
			return
		},
	}
)
