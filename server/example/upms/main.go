// @Title
// @Description
// @Author  Wangwengang  2023/12/16 14:33
// @Update  Wangwengang  2023/12/16 14:33
package main

import (
	"github.com/smallnest/rpcx/server"
	"github.com/wwengg/simple/core/sconfig"
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/snet/http"
	"github.com/wwengg/simple/core/srpc"
	"github.com/wwengg/simple/core/store"
	"github.com/wwengg/simple/server/example/upms/global"
	"github.com/wwengg/simple/server/example/upms/model"
	"github.com/wwengg/simple/server/example/upms/service"
)

func main() {
	// 读取配置文件
	sconfig.Viper(&global.CONFIG)

	// 初始化日志
	global.Log = slog.NewZapLog(&global.CONFIG.Slog)

	global.Log.Info("获取到的配置信息", slog.Any("CONFIG", global.CONFIG))

	// 初始化RPCClient
	global.RPC = srpc.NewSRPCClients(&global.CONFIG.RPC)

	// 初始化DBList
	global.DBList = store.DBList(&global.CONFIG.DBList)

	// upms数据库
	global.DBUpms = global.DBList["upms"]

	// 创建初始化upms的数据库表
	global.DBUpms.AutoMigrate(
		model.SimpleUser{},
	)

	// 初始化rpcx服务
	s := server.NewServer()

	// 服务注册中心
	srpc.AddRegistryPlugin(s, global.CONFIG.RPC, global.CONFIG.RpcService)

	s.RegisterName("User", new(service.User), "")

	// 执行上次程序退出未处理完的数据
	initServer()

	// go 协程启动rpc服务 开始接客
	go global.Log.Error(s.Serve("tcp", global.CONFIG.RpcService.ServiceAddr).Error())

	// 优雅的停止服务
	http.InitServer("localhost:4399", nil) // 实际进程，无实际作用，只为停止服务的时候有时间处理未处理完的数据，4399端口如占用自行修改成7k7k(暴露年龄了

	// 检测到进程退出执行以下代码，优雅的退出服务，处理未处理完的数据（以下所有代码对windows无效）
	global.Log.Infof("检测到程序退出，储存数据，下次启动再继续执行")

}

func initServer() {

}
