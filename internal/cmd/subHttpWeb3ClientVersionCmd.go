package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	HttpWeb3ClientVersion = gcmd.Command{
		Name:  "httpWeb3ClientVersion",
		Usage: "httpWeb3ClientVersion",
		Brief: "http web3ClientVersion",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			msg := `{ 
				"jsonrpc": "2.0",
				"method": "web3_clientVersion",
				"params": [],
				"id": 1337
				}`

			rsp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(msg)))
			assert(err)
			ret, err := ioutil.ReadAll(rsp.Body)
			assert(err)
			fmt.Println("http web3_clientVersion \r\n", string(ret))
			return
		},
	}
)
