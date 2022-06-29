package eth

import (
	"context"
	"fmt"
	_ "fmt"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

type SimpleTx struct {
	From     common.Address `json:"from"`
	To       common.Address `json:"to"`
	Gas      uint64         `json:"gas"`
	GasPrice *big.Int       `json:"gasPrice"`
	Value    *big.Int       `json:"value"`
	Data     []byte         `json:"data"`
	Nonce    uint64         `json:"nonce"`
}

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

func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

type Client struct {
	RpcClient *rpc.Client
}

func Dial(url string) (*Client, error) {
	rpcClient, err := rpc.Dial(url)
	if err != nil {
		return nil, err
	}

	var client = Client{RpcClient: rpcClient}
	return &client, nil
}

func (c *Client) Web3ClientVersion(ctx context.Context) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "web3_clientVersion")
	return result, err
}

func (c *Client) Web3Sha3(ctx context.Context, data string) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "web3_sha3", data)
	return result, err
}

func (c *Client) NetVersion(ctx context.Context) (*big.Int, error) {
	var netVersion big.Int
	err := c.RpcClient.CallContext(ctx, &netVersion, "net_version")
	return &netVersion, err
}

func (c *Client) NetListening(ctx context.Context) (bool, error) {
	var result bool
	err := c.RpcClient.CallContext(ctx, &result, "net_listening")
	return result, err
}

func (c *Client) EthProtocolVersion(ctx context.Context) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "eth_protocolVersion")
	return result, err
}

func (c *Client) NetPeerCount(ctx context.Context) (uint, error) {
	var result uint
	err := c.RpcClient.CallContext(ctx, &result, "net_peerCount")
	return result, err
}

func (c *Client) EthCoinbase(ctx context.Context) (common.Address, error) {
	var result common.Address
	err := c.RpcClient.CallContext(ctx, &result, "eth_coinbase")
	return result, err
}

func (c *Client) EthMining(ctx context.Context) (bool, error) {
	var result bool
	err := c.RpcClient.CallContext(ctx, &result, "eth_mining")
	return result, err
}

func (c *Client) EthHashrate(ctx context.Context) (uint64, error) {
	var result hexutil.Uint64
	err := c.RpcClient.CallContext(ctx, &result, "eth_hashrate")
	return uint64(result), err
}

func (c *Client) EthAccounts(ctx context.Context) ([]common.Address, error) {
	var result []common.Address
	err := c.RpcClient.CallContext(ctx, &result, "eth_accounts")
	return result, err
}

func (c *Client) EthGetBalance(ctx context.Context, account common.Address, blockLocator interface{}) (*big.Int, error) {
	var result hexutil.Big
	err := c.RpcClient.CallContext(ctx, &result, "eth_getBalance", account, blockLocator)
	return (*big.Int)(&result), err
}

func (c *Client) EthGetTransactionCount(ctx context.Context, account common.Address, blockLocator interface{}) (uint64, error) {
	var result hexutil.Uint64
	err := c.RpcClient.CallContext(ctx, &result, "eth_getTransactionCount", account, blockLocator)
	return uint64(result), err
}

func (c *Client) EthGetBlockTransactionCountByNumber(ctx context.Context, blockLocator interface{}) (uint, error) {
	var result hexutil.Uint
	err := c.RpcClient.CallContext(ctx, &result, "eth_getBlockTransactionCountByNumber", blockLocator)
	return uint(result), err
}

func (c *Client) EthGetUncleCountByBlockHash(ctx context.Context, blockHash common.Hash) (uint, error) {
	var result hexutil.Uint
	err := c.RpcClient.CallContext(ctx, &result, "eth_getUncleCountByBlockHash", blockHash)
	return uint(result), err
}

func (c *Client) EthGetUncleCountByBlockNumber(ctx context.Context, blockLocator interface{}) (uint, error) {
	var result hexutil.Uint
	err := c.RpcClient.CallContext(ctx, &result, "eth_getUncleCountByBlockNumber", blockLocator)
	return uint(result), err
}

func (c *Client) EthGetCode(ctx context.Context, account common.Address, blockLocator interface{}) ([]byte, error) {
	var result hexutil.Bytes
	err := c.RpcClient.CallContext(ctx, &result, "eth_getCode", account, blockLocator)
	return result, err
}

func (c *Client) EthSign(ctx context.Context, account common.Address, data []byte) ([]byte, error) {
	var result hexutil.Bytes
	err := c.RpcClient.CallContext(ctx, &result, "eth_sign", account, data)
	return result, err
}

func (c *Client) EthSendTransaction(ctx context.Context, msg interface{}) (common.Hash, error) {
	var result common.Hash
	err := c.RpcClient.CallContext(ctx, &result, "eth_sendTransaction", msg)
	return result, err
}

func (c *Client) EthSendRawTransaction(ctx context.Context, rawtx []byte) (common.Hash, error) {
	var result common.Hash
	//hex := common.ToHex(rawtx)
	hex := "0x" + common.Bytes2Hex(rawtx)
	//hex := common.Bytes2Hex(rawtx)
	fmt.Println("hex: ", hex)
	err := c.RpcClient.CallContext(ctx, &result, "eth_sendRawTransaction", hex)
	return result, err
}

func (c *Client) EthCall(ctx context.Context, msg interface{}) ([]byte, error) {
	var result hexutil.Bytes
	err := c.RpcClient.CallContext(ctx, &result, "eth_call", msg)
	return result, err
}

func (c *Client) EthEstimateGas(ctx context.Context, msg interface{}) (uint64, error) {
	var result hexutil.Uint64
	err := c.RpcClient.CallContext(ctx, &result, "eth_estimateGas", msg)
	return uint64(result), err
}

func (c *Client) EthGetTransactionReceipt(ctx context.Context, hash common.Hash) (*types.Receipt, error) {
	var result types.Receipt
	err := c.RpcClient.CallContext(ctx, &result, "eth_getTransactionReceipt", hash)
	return &result, err
}

func (c *Client) EthNewBlockFilter(ctx context.Context) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "eth_newBlockFilter")
	return result, err
}

func (c *Client) EthGetFilterChanges(ctx context.Context, filterId string) ([]common.Hash, error) {
	var result []common.Hash
	err := c.RpcClient.CallContext(ctx, &result, "eth_getFilterChanges", filterId)
	return result, err
}

func (c *Client) EthGetFilterLogs(ctx context.Context, filterId uint64) (*types.Log, error) {
	var result types.Log
	err := c.RpcClient.CallContext(ctx, &result, "eth_getFilterLogs", filterId)
	return &result, err
}

func (c *Client) EthNewPendingTransactionFilter(ctx context.Context) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "eth_newPendingTransactionFilter")
	return result, err
}

func (c *Client) EthNewFilter(ctx context.Context, filterOpts interface{}) (string, error) {
	var result string
	err := c.RpcClient.CallContext(ctx, &result, "eth_newFilter", filterOpts)
	return result, err
}

func (c *Client) EthGetLogFilterChanges(ctx context.Context, filterId string) ([]types.Log, error) {
	var result []types.Log
	err := c.RpcClient.CallContext(ctx, &result, "eth_getFilterChanges", filterId)
	return result, err
}

func (c *Client) SubscribeNewPendingTransactions(ctx context.Context, ch chan<- *common.Hash) (ethereum.Subscription, error) {
	return c.RpcClient.EthSubscribe(ctx, ch, "newPendingTransactions")
}
