package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogf/gf/v2/os/gcmd"
)

type RpcRequest struct {
	JsonRpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	Id      int64         `json:"id"`
}

type RpcResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	Id      int64       `json:"id"`
	Result  interface{} `json:"result"`
	Error   interface{} `json:"error"`
}

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
			rpcRequest := RpcRequest{
				JsonRpc: "2.0",
				Method:  "web3_clientVersion",
				Params:  []interface{}{}, //type interface{}, {} is null, [] is array
				Id:      time.Now().Unix(),
			}
			payload, err := json.Marshal(rpcRequest)
			assert(err)
			fmt.Println("payload: ", payload)
			rsp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(payload)))
			assert(err)
			ret, err := ioutil.ReadAll(rsp.Body)
			assert(err)
			fmt.Println("http web3_clientVersion \r\n", string(ret))
			var rpcResponse RpcResponse
			err = json.Unmarshal(ret, &rpcResponse)
			assert(err)
			fmt.Println("http web3_clientVersion \r\n", rpcResponse)
			return
		},
	}
)
