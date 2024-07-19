package sbus

import (
	"context"
	"net"
)

type SConnection interface {
	// Start the connection, make the current connection start working
	// (启动连接，让当前连接开始工作)
	Start()
	// Stop the connection and end the current connection state
	// (停止连接，结束当前连接状态)
	Stop()

	// Returns ctx, used by user-defined go routines to obtain connection exit status
	// (返回ctx，用于用户自定义的go程获取连接退出状态)
	Context() context.Context

	GetConnID() uint64            // Get the current connection ID (获取当前连接ID)
	GetConnIdStr() string         // Get the current connection ID for string (获取当前字符串连接ID)
	GetTaskHandler() STaskHandler // Get the message handler (获取消息处理器)
	RemoteAddr() net.Addr         // Get the remote address information of the connection (获取链接远程地址信息)
	LocalAddr() net.Addr          // Get the local address information of the connection (获取链接本地地址信息)
	LocalAddrString() string      // Get the local address information of the connection as a string
	RemoteAddrString() string     // Get the remote address information of the connection as a string

	SendData(data []byte) error // Send data to the message queue to be sent to the remote TCP client later
	SendMsg(msgID uint32, data []byte) error

	SetProperty(key string, value string)   // Set connection property
	GetProperty(key string) (string, error) // Get connection property
	RemoveProperty(key string)              // Remove connection property
	IsAlive() bool                          // Check if the current connection is alive(判断当前连接是否存活)
	SetHeartBeat(checker SHeartbeatChecker) // Set the heartbeat detector (设置心跳检测器)

	//AddCloseCallback(handler, key interface{}, callback func()) // Add a close callback function (添加关闭回调函数)
	//RemoveCloseCallback(handler, key interface{})               // Remove a close callback function (删除关闭回调函数)
	//InvokeCloseCallbacks()                                      // Trigger the close callback function (触发关闭回调函数，独立协程完成)
}
