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

	//use eth tools
	cmd.Main.AddCommand(&cmd.EthToolsDial)
	cmd.Main.AddCommand(&cmd.EthToolsGetBalance)
	cmd.Main.AddCommand(&cmd.EthToolsSendTransaction)
	cmd.Main.AddCommand(&cmd.EthToolsSendRawTransaction)
	cmd.Main.AddCommand(&cmd.EthToolsDeployAccessContract1)
	cmd.Main.AddCommand(&cmd.EthToolsDeployAccessContract2GoVersion)
	cmd.Main.AddCommand(&cmd.EthToolsBlockMonitorPull)
	cmd.Main.AddCommand(&cmd.EthToolsBlockMonitorPush)
	cmd.Main.AddCommand(&cmd.EthToolsBlockPendingMonitorPull)
	cmd.Main.AddCommand(&cmd.EthToolsBlockPendingMonitorPush)
	cmd.Main.AddCommand(&cmd.EthToolsLogsMonitorPull)

	// private key, public key & account
	cmd.Main.AddCommand(&cmd.GenerateKey)
	cmd.Main.AddCommand(&cmd.GenerateKeyStoreFile)
	cmd.Main.Run(gctx.New())
}
