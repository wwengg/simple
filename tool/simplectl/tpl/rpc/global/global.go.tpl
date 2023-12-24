/*
{{ .Copyright }}
{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
*/
package global

import (
	"github.com/wwengg/simple/core/slog"
	"github.com/wwengg/simple/core/srpc"
	"github.com/wwengg/simple/core/store"
	"gorm.io/gorm"
)

var (
	CONFIG *Config
	LOG    slog.Slog
	SRPC   srpc.SRPC
	DBList map[string]*gorm.DB
)


func InitSlog() {
	// 初始化日志
	LOG = slog.NewZapLog(&CONFIG.Slog)
}

func InitSRPC() {
	// 初始化SRPC
	SRPC = srpc.NewSRPCClients(&CONFIG.RPC)
}

func InitDB() {
	// 初始化DBList
	DBList = store.DBList(&CONFIG.DBList)

	// 创建初始化数据库表
	//DBUpms.AutoMigrate(
	//	model.ServerInfo{},
	//)
}