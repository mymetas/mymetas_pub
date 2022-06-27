package cmd

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/os/gcmd"
)

func generateKey() {
	fmt.Println("new account demo")

	prvKey, err := crypto.GenerateKey()
	eth.Assert(err)

	fmt.Println("prevKey is:", prvKey)
	prvKeyBytes := crypto.FromECDSA(prvKey)
	fmt.Println("private key: ", hexutil.Encode(prvKeyBytes))

	pubKey := prvKey.Public()
	pubKeyECDSA := pubKey.(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	fmt.Println("public key: ", hexutil.Encode(pubKeyBytes))

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println("address: ", address.Hex())
}

var (
	GenerateKey = gcmd.Command{
		Name:  "generateKey",
		Usage: "generateKey",
		Brief: " generate key & keystore ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			generateKey()
			return
		},
	}
)
