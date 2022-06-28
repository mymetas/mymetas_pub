package cmd

import (
	"context"
	"fmt"
	"mymetas_pub/internal/service/eth"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/google/uuid"
)

// init
var (
	fileName = "keyStore.json"
	//tempDir  = gfile.TempDir("gfile_example_content")
	tempDir  = gfile.Temp("key_store")
	tempFile = gfile.Join(tempDir, fileName)

	password = "password"
)

func encodeKeyStore(keyHex, auth string) []byte {
	prvKey, err := crypto.HexToECDSA(keyHex[2:])
	eth.Assert(err)

	PrvKeyToPublicAndAddress(prvKey)

	id, err := uuid.NewRandom()
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(prvKey.PublicKey),
		PrivateKey: prvKey,
	}
	json, err := keystore.EncryptKey(key, auth, keystore.StandardScryptN, keystore.StandardScryptP)
	eth.Assert(err)
	return json
}

func reloadKeyStore() string {
	json := gfile.GetContents(tempFile)
	key, err := keystore.DecryptKey([]byte(json), password)
	eth.Assert(err)
	return hexutil.Encode(crypto.FromECDSA(key.PrivateKey))
}

func saveKeyStore(keyStore []byte) {
	// write contents
	gfile.PutBytes(tempFile, keyStore)

	// read contents
	fmt.Println(gfile.GetContents(tempFile))
}

func generateKeyStoreFile() {
	privateKey := "0xa065cbef250c56523fd8e117122099a71ccdcd4261960bc3fd620d8c119d8f3d"
	fmt.Println("\r\nprivate key: ------", privateKey)

	keyStoreByte := encodeKeyStore(privateKey, password)
	//fmt.Println("encoded: ", string(keyStoreByte))

	saveKeyStore(keyStoreByte)

	decoded := reloadKeyStore()
	fmt.Println("\r\nreloadKeyStore decoded private key: ------", decoded)
}

var (
	GenerateKeyStoreFile = gcmd.Command{
		Name:  "generateKeyStoreFile",
		Usage: "generateKeyStoreFile",
		Brief: " generate keyStore, write & get file ",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			generateKeyStoreFile()
			return
		},
	}
)
