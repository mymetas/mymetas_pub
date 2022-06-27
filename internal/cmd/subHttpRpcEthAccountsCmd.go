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
	HttpRpcEthAccounts = gcmd.Command{
		Name:  "httpRpcEthAccounts",
		Usage: "httpRpcEthAccounts",
		Brief: "http Rpc ethAccounts",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			rpcRequest := eth.RpcRequest{
				JsonRpc: "2.0",
				Method:  "eth_accounts",
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
			fmt.Println("http web3_eth_accounts response: \r\n", string(ret))
			var rpcResponse eth.RpcResponse
			err = json.Unmarshal(ret, &rpcResponse)
			eth.Assert(err)
			fmt.Println("http web3_eth_accounts response unmarshal\r\n", rpcResponse)
			return
		},
	}
)
