package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/gogf/gf/v2/os/gcmd"
)

func ethToolsDialWeb3ClientVersion() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)
	version, err := client.Web3ClientVersion(context.Background())
	eth.Assert(err)
	fmt.Println(version)
}

func ethToolsDialWeb3Sha3() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	bytes := []byte("this is demo")
	//notice  byte  2  hex, must add 0x
	data := "0x" + common.Bytes2Hex(bytes)

	hash, err := client.Web3Sha3(context.Background(), data)
	eth.Assert(err)
	fmt.Println(hash)
}

var (
	EthToolsDial = gcmd.Command{
		Name:  "ethToolsDial",
		Usage: "ethToolsDial",
		Brief: "eth tools dial  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			ethToolsDialWeb3ClientVersion()
			ethToolsDialWeb3Sha3()
			return
		},
	}
)
