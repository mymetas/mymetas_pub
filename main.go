package main

import (
	_ "mymetas_pub/internal/packed"

	"mymetas_pub/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.AddCommand(&cmd.HttpWeb3ClientVersion)
	cmd.Main.AddCommand(&cmd.HttpRpcWeb3ClientVersion)
	cmd.Main.AddCommand(&cmd.HttpRpcEthAccounts)
	cmd.Main.AddCommand(&cmd.EthRpcDial)
	cmd.Main.AddCommand(&cmd.EthClient)
	cmd.Main.Run(gctx.New())
}
