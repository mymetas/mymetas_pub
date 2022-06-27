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

var (
	HttpRpcEthAccounts = gcmd.Command{
		Name:  "httpRpcEthAccounts",
		Usage: "httpRpcEthAccounts",
		Brief: "http Rpc ethAccounts",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			rpcRequest := RpcRequest{
				JsonRpc: "2.0",
				Method:  "eth_accounts",
				Params:  []interface{}{}, //type interface{}, {} is null, [] is array
				Id:      time.Now().Unix(),
			}
			payload, err := json.Marshal(rpcRequest)
			assert(err)
			fmt.Println("payload: ", string(payload))
			rsp, err := http.Post("http://localhost:8545", "application/json", bytes.NewBuffer([]byte(payload)))
			assert(err)
			ret, err := ioutil.ReadAll(rsp.Body)
			assert(err)
			fmt.Println("http web3_eth_accounts response: \r\n", string(ret))
			var rpcResponse RpcResponse
			err = json.Unmarshal(ret, &rpcResponse)
			assert(err)
			fmt.Println("http web3_eth_accounts response unmarshal\r\n", rpcResponse)
			return
		},
	}
)
