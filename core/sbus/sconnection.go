package sbus

import (
	"context"
	"net"
	"time"
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

	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	// Write can be made to time out and return an error after a fixed
	// time limit; see SetDeadline and SetWriteDeadline.
	Write(b []byte) (n int, err error)

	// Close closes the connection.
	// Any blocked Read or Write operations will be unblocked and return errors.
	Close() error

	// SetDeadline sets the read and write deadlines associated
	// with the connection. It is equivalent to calling both
	// SetReadDeadline and SetWriteDeadline.
	//
	// A deadline is an absolute time after which I/O operations
	// fail instead of blocking. The deadline applies to all future
	// and pending I/O, not just the immediately following call to
	// Read or Write. After a deadline has been exceeded, the
	// connection can be refreshed by setting a deadline in the future.
	//
	// If the deadline is exceeded a call to Read or Write or to other
	// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
	// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
	// The error's Timeout method will return true, but note that there
	// are other possible errors for which the Timeout method will
	// return true even if the deadline has not been exceeded.
	//
	// An idle timeout can be implemented by repeatedly extending
	// the deadline after successful Read or Write calls.
	//
	// A zero value for t means I/O operations will not time out.
	SetDeadline(t time.Time) error

	// SetReadDeadline sets the deadline for future Read calls
	// and any currently-blocked Read call.
	// A zero value for t means Read will not time out.
	SetReadDeadline(t time.Time) error

	// SetWriteDeadline sets the deadline for future Write calls
	// and any currently-blocked Write call.
	// Even if write times out, it may return n > 0, indicating that
	// some of the data was successfully written.
	// A zero value for t means Write will not time out.
	SetWriteDeadline(t time.Time) error
}
