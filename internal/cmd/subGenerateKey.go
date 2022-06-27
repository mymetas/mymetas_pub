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

func prvKeyToPublicAndAddress(prvKey *ecdsa.PrivateKey) string {
	fmt.Println("prevKey is:", prvKey)
	prvKeyBytes := crypto.FromECDSA(prvKey)
	generatePrivateKey := hexutil.Encode(prvKeyBytes)
	fmt.Println("private key: ", generatePrivateKey)

	pubKey := prvKey.Public()
	pubKeyECDSA := pubKey.(*ecdsa.PublicKey)
	pubKeyBytes := crypto.FromECDSAPub(pubKeyECDSA)
	fmt.Println("public key: ", hexutil.Encode(pubKeyBytes))

	address := crypto.PubkeyToAddress(*pubKeyECDSA)
	fmt.Println("address: ", address.Hex())

	return generatePrivateKey
}

func generateKey() string {
	fmt.Println("new account demo")

	prvKey, err := crypto.GenerateKey()
	eth.Assert(err)

	return prvKeyToPublicAndAddress(prvKey)
}

func importPrivateKey(hexPrivate string) {
	fmt.Println("\r\nimportPrivateKey: ", hexPrivate)
	prvKey, err := crypto.HexToECDSA(hexPrivate[2:])
	eth.Assert(err)

	privateKey := prvKeyToPublicAndAddress(prvKey)
	fmt.Println("\r\nexport PrivateKey: ", privateKey)

	if privateKey != hexPrivate {
		panic("private key prvKeyToPublicAndAddress error")
	}
}

var (
	GenerateKey = gcmd.Command{
		Name:  "generateKey",
		Usage: "generateKey",
		Brief: " generate key & keystore ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			privateKey := generateKey()
			importPrivateKey(privateKey)
			return
		},
	}
)
