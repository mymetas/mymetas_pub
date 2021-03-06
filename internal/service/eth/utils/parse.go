package eth

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func EncodeUintToHex(i uint64) string {
	return fmt.Sprintf("0x%x", i)
}

func ParseBigInt(str string) *big.Int {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	num := new(big.Int)
	num.SetString(str, 16)
	return num
}

func ParseUint64orHex(str string) (uint64, error) {
	base := 10
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
		base = 16
	}
	return strconv.ParseUint(str, base, 64)
}

func EncodeToHex(b []byte) string {
	return "0x" + hex.EncodeToString(b)
}

func ParseHexBytes(str string) ([]byte, error) {
	if !strings.HasPrefix(str, "0x") {
		return nil, fmt.Errorf("it does not have 0x prefix")
	}
	buf, err := hex.DecodeString(str[2:])
	if err != nil {
		return nil, err
	}
	return buf, nil
}
