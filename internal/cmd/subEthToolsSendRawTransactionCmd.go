package cmd

import (
	"bytes"
	"context"
	"fmt"
	"math/big"
	"mymetas_pub/internal/service/eth"
	web3 "mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func ethClientNetworkId() *big.Int {
	client, err := ethclient.Dial("http://localhost:8545")
	eth.Assert(err)

	networkId, err := client.NetworkID(context.Background())
	eth.Assert(err)
	return networkId
}

func ethToolsSendRawTransaction1() {
	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	// address, err := client.EthAccounts(ctx)
	// eth.Assert(err)
	// DisplayBalance(ctx, *client, address)

	networkId := ethClientNetworkId()

	//0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d
	//is ganache -d address[0]
	credential, err := eth.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	eth.Assert(err)
	fmt.Println("from address: ", credential.Address.Hex())

	to := common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
	fmt.Println("to address: ", to.Hex())

	nonce, err := client.EthGetTransactionCount(ctx, credential.Address, "pending")
	eth.Assert(err)
	fmt.Println("nonce: ", nonce)

	tx := types.NewTransaction(
		nonce,
		to,
		big.NewInt(1000000),
		uint64(21000),
		big.NewInt(2e9),
		nil,
	)

	signedTx, err := credential.SignTx(tx, networkId)
	eth.Assert(err)
	fmt.Println("signedTx: ", signedTx)

	buf := new(bytes.Buffer)
	err = signedTx.EncodeRLP(buf)
	eth.Assert(err)

	fmt.Println("buf is :", buf.String())
	txid, err := client.EthSendRawTransaction(ctx, buf.Bytes())
	eth.Assert(err)
	fmt.Println("raw tx id: ", txid.Hex())

	//DisplayBalance(ctx, *client, address)
}

func ethToolsSendRawTransaction() {
	credential, err := eth.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	eth.Assert(err)
	fmt.Println("from address: ", credential.Address.Hex())

	to := common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")
	fmt.Println("to address: ", to.Hex())

	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)

	ctx := context.Background()

	chainid, err := client.NetVersion(ctx)
	eth.Assert(err)
	fmt.Println("chainid: ", chainid)

	nonce, err := client.EthGetTransactionCount(ctx, credential.Address, "pending")
	eth.Assert(err)
	fmt.Println("nonce: ", nonce)

	tx := types.NewTransaction(
		nonce,
		to,
		big.NewInt(1000000),
		uint64(21000),
		big.NewInt(2e9),
		nil,
	)
	signedTx, err := credential.SignTx(tx, chainid)
	eth.Assert(err)
	fmt.Println("signedTx: ", signedTx)

	buf := new(bytes.Buffer)
	err = signedTx.EncodeRLP(buf)
	eth.Assert(err)
	txid, err := client.EthSendRawTransaction(ctx, buf.Bytes())
	eth.Assert(err)
	fmt.Println("raw tx id: ", txid.Hex())

	balance, err := client.EthGetBalance(ctx, to, "latest")
	eth.Assert(err)
	fmt.Println("balance received: ", balance)
}

func ethToolsSendRawTransaction2() {
	credential, err := eth.HexToCredential("0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d")
	eth.Assert(err)
	fmt.Println("from address: ", credential.Address.Hex())

	to := common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0")

	client, err := eth.Dial("http://localhost:8545")
	eth.Assert(err)
	ctx := context.Background()
	address, err := client.EthAccounts(ctx)
	eth.Assert(err)
	DisplayBalance(ctx, *client, address)

	ethclienttmp, err := ethclient.Dial("http://localhost:8545")
	eth.Assert(err)

	nonce, err := ethclienttmp.NonceAt(context.Background(), address[0], nil)
	eth.Assert(err)

	amount := big.NewInt(1000000)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(10)
	data := []byte{}

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	// fmt.Println(tx)
	chainId := ethClientNetworkId()
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), credential.PrivateKey)
	eth.Assert(err)
	fmt.Println("signTx", signedTx)
	fmt.Println("privateKey:", credential.PrivateKey)

	serializedTx, err := rlp.EncodeToBytes(signedTx)
	eth.Assert(err)
	fmt.Printf("serializedTx 0x%x\n", serializedTx)

	//err = e.c.Call("eth_sendRawTransaction", &hash, fmt.Sprintf("0x%x", serializedTx))
	hash, err := client.EthSendRawTransaction2(ctx, fmt.Sprintf("0x%x", serializedTx))
	eth.Assert(err)
	fmt.Println("sendRawTransaction hash : ", hash)
}

func ethToolsSendRawTransaction3() {
	// change to your rpc provider
	var rpcProvider = "http://localhost:8545"

	web3, err := web3.NewWeb3(rpcProvider)

	if err != nil {
		panic(err)
	}
	privateString := "4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d"
	web3.Eth.SetAccount(privateString)
	// set default account by private key
	privateKey := privateString
	kovanChainId := int64(1337)
	if err := web3.Eth.SetAccount(privateKey); err != nil {
		panic(err)
	}
	web3.Eth.SetChainId(kovanChainId)

	txHash, err := web3.Eth.SyncSendRawTransaction(
		common.HexToAddress("0xFFcf8FDEE72ac11b5c542428B35EEF5769C409f0"),
		big.NewInt(1000),
		uint64(21000),
		web3.Utils.ToGWei(1),
		nil,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Send approve tx hash %v\n", txHash)
}

var (
	EthToolsSendRawTransaction = gcmd.Command{
		Name:  "ethToolsSendRawTransaction",
		Usage: "ethToolsSendRawTransaction",
		Brief: "ethTools SendRawTransaction  ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			client, err := eth.Dial("http://localhost:8545")
			eth.Assert(err)
			address, err := client.EthAccounts(ctx)
			eth.Assert(err)
			DisplayBalance(ctx, *client, address)

			ethToolsSendRawTransaction3()

			DisplayBalance(ctx, *client, address)
			return
		},
	}
)
