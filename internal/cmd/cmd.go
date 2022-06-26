package cmd

import (
	"context"
	"fmt"

	"mymetas_pub/internal/controller"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
				)
			})
			s.Run()
			return nil
		},
	}
	HttpWeb3ClientVersion = gcmd.Command{
		Name:  "httpWeb3ClientVersion",
		Usage: "httpWeb3ClientVersion",
		Brief: "http web3ClientVersion",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			fmt.Println("http web3Clientversion")
			return
		},
	}
)
