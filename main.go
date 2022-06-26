package main

import (
	_ "mymetas_pub/internal/packed"

	"mymetas_pub/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.AddCommand(&cmd.HttpWeb3ClientVersion)
	cmd.Main.Run(gctx.New())
}
