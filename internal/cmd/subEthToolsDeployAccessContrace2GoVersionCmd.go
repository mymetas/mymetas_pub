package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"mymetas_pub/internal/service/eth"
	testerc20token "mymetas_pub/internal/service/eth/contract/wrapper/testerc20token"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/os/gcmd"
)

func deployHappyTokenContract2GoVersion() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	credential, err := eth.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	eth.Assert(err)

	tokenSupply := big.NewInt(1000000)
	tokenName := "HAPPY TOKEN"
	tokenDecimals := uint8(0)
	tokenSymbol := "HAPY"

	txOpts := bind.NewKeyedTransactor(credential.PrivateKey)

	//transaction type not supported,     error..................... gasPrice fix it
	//auth := bind.NewKeyedTransactor(credential.PrivateKey)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	txOpts.GasPrice = gasPrice

	address, tx, inst, err := testerc20token.DeployTesterc20token(txOpts, client, tokenSupply, tokenName, tokenDecimals, tokenSymbol)
	eth.Assert(err)
	fmt.Println("deployed at: ", address.Hex())
	fmt.Println("txid: ", tx.Hash().Hex())
	_ = inst

	fmt.Println("save deployed address...")
	err = ioutil.WriteFile(HappyTokenAddrPath, []byte(address.Hex()), 0644)
	eth.Assert(err)

	fmt.Println("done.")
}

func accessContractTransfer2GoVersion() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	addrHexBytes, err := ioutil.ReadFile(HappyTokenAddrPath)
	eth.Assert(err)
	contractAddress := common.HexToAddress(string(addrHexBytes))
	fmt.Println("contract address: ", contractAddress.Hex())

	inst, err := testerc20token.NewTesterc20token(contractAddress, client)
	eth.Assert(err)
	//fmt.Println("inst: ",inst)

	credential, err := eth.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	eth.Assert(err)

	txOpts := bind.NewKeyedTransactor(credential.PrivateKey)
	//sty modify, fix error................ transaction type not supported
	gasPrice, err := client.SuggestGasPrice(context.Background())
	txOpts.GasPrice = gasPrice

	toAddress := common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
	amount := big.NewInt(100)

	tx, err := inst.Transfer(txOpts, toAddress, amount)
	eth.Assert(err)
	fmt.Println("txid: ", tx.Hash().Hex())

	callOpts := &bind.CallOpts{
		From: credential.Address,
	}
	balance, err := inst.BalanceOf(callOpts, toAddress)
	eth.Assert(err)
	fmt.Println("balance: ", balance)
}

func ethToolsDeployAccessContract2GoVersion() {
	deployHappyTokenContract2GoVersion()
	accessContractTransfer2GoVersion()
}

var (
	EthToolsDeployAccessContract2GoVersion = gcmd.Command{
		Name:  "ethToolsDeployAccessContract2GoVersion",
		Usage: "ethToolsDeployAccessContract2GoVersion",
		Brief: "ethTools Deploy && AccessContract 2 Go Version  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsDeployAccessContract2GoVersion()
			return
		},
	}
)
