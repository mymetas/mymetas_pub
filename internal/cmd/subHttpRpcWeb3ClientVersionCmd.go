package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mymetas_pub/internal/service/eth"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	HttpRpcWeb3ClientVersion = gcmd.Command{
		Name:  "httpRpcWeb3ClientVersion",
		Usage: "httpRpcWeb3ClientVersion",
		Brief: "http Rpc web3ClientVersion",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// msg := `{
			// 	"jsonrpc": "2.0",
			// 	"method": "web3_clientVersion",
			// 	"params": [],
			// 	"id": 1337
			// 	}`
			rpcRequest := eth.RpcRequest{
				JsonRpc: "2.0",
				Method:  "web3_clientVersion",
				Params:  []interface{}{}, //type interface{}, {} is null, [] is array
				Id:      time.Now().Unix(),
			}
			payload, err := json.Marshal(rpcRequest)
			eth.Assert(err)
			fmt.Println("payload: ", string(payload))
			rsp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(payload)))
			eth.Assert(err)
			ret, err := ioutil.ReadAll(rsp.Body)
			eth.Assert(err)
			fmt.Println("http web3_clientVersion response: \r\n", string(ret))
			var rpcResponse eth.RpcResponse
			err = json.Unmarshal(ret, &rpcResponse)
			eth.Assert(err)
			fmt.Println("http web3_clientVersion response unmarshal\r\n", rpcResponse)
			return
		},
	}
)
