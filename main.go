package main

import (
	_ "mymetas_pub/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"mymetas_pub/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
