package cmd

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/gogf/gf/v2/os/gcmd"
)

func rpcDialWeb3ClientVersion() {
	client, _ := rpc.Dial("http://localhost:8545")
	var version string
	_ = client.Call(&version, "web3_clientVersion")
	fmt.Println(version)
}

func rpcDialWeb3Sha3() {
	client, _ := rpc.Dial("http://localhost:8545")
	bytes := []byte("this is demo")

	//notice  byte  2  hex, must add 0x
	data := "0x" + common.Bytes2Hex(bytes)
	fmt.Println("data is: ", string(data))

	var hash string
	_ = client.Call(&hash, "web3_sha3", data)
	fmt.Println("hash: ", hash)
}

var (
	EthRpcDial = gcmd.Command{
		Name:  "ethRpcDial",
		Usage: "ethRpcDial",
		Brief: "eth rpc dial  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			rpcDialWeb3ClientVersion()
			rpcDialWeb3Sha3()
			return
		},
	}
)
