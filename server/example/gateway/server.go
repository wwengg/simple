// @Title
// @Description
// @Author  Wangwengang  2023/12/10 13:44
// @Update  Wangwengang  2023/12/10 13:44
package main

import (
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/snet/http"
	"github.com/wwengg/simple/core/srpc"
	"github.com/wwengg/simple/server"
	"github.com/wwengg/simple/server/example/gateway/global"
	"github.com/wwengg/simple/server/example/gateway/middleware"
	"github.com/wwengg/simple/server/example/gateway/router"
)

func main() {
	// 初始化配置文件
	sconfig.Viper("./server/example/gateway/config.yaml")

	// 初始化日志
	global.SLog = slog.NewZapLog(&sconfig.S_CONF.Slog)

	// 初始化SRPC
	global.SRPC = srpc.NewSRPCClients(&sconfig.S_CONF.RPC)

	// 初始化gin
	ginEngine := http.NewGinEngine(&sconfig.S_CONF)

	// 配置路由，中间件
	publicGroup := ginEngine.GetPublicRouterGroup()
	privateGroup := ginEngine.GetPrivateRouterGroup()
	publicGroup.Use(middleware.BaseHandler())
	{
		router.InitSRPCRouter(publicGroup)
	}
	privateGroup.Use(middleware.BaseHandler())
	{
		router.InitSRPCRouter(privateGroup)
	}

	srv := server.NewGateway(&sconfig.S_CONF.Gateway, ginEngine)

	srv.Start()
}
