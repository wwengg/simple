// @Title
// @Description
// @Author  Wangwengang  2023/12/16 00:50
// @Update  Wangwengang  2023/12/16 00:50
package im

import (
	"context"
	"golang.org/x/net/websocket"
	"net"
)

type SConnection interface {
	Start() // 开始工作
	Stop()  // 结束连接

	Context() context.Context // 获取上下文管理器

	// 获取原始链接
	GetConnection() net.Conn
	GetWsConn() *websocket.Conn

	Send(data []byte) error  // 直接发数据包
	SendToQueue(data []byte) // 通过消息队列发送数据包

	IsAlive() bool // 判断当前连接是否存活
	SetHeartBeat()
}
