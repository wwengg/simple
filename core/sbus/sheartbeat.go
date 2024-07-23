package sbus

type SHeartbeatChecker interface {
	SetOnRemoteNotAlive(OnRemoteNotAlive)
	SetHeartbeatMsgFunc(HeartBeatMsgFunc)
	SetHeartbeatFunc(HeartBeatFunc)
	BindRouter(uint32, SRouter)
	Start()
	Stop()
	SendHeartBeatMsg() error
	BindConn(connection SConnection)
	Clone() SHeartbeatChecker
	MsgID() uint32
	Router() SRouter
}

// User-defined method for handling heartbeat detection messages
// (用户自定义的心跳检测消息处理方法)
type HeartBeatMsgFunc func(connection SConnection) []byte

// HeartBeatFunc User-defined heartbeat function
// (用户自定义心跳函数)
type HeartBeatFunc func(connection SConnection) error

// OnRemoteNotAlive User-defined method for handling remote connections that are not alive
// 用户自定义的远程连接不存活时的处理方法
type OnRemoteNotAlive func(connection SConnection)

type HeartBeatOption struct {
	MakeMsg          HeartBeatMsgFunc // User-defined method for handling heartbeat detection messages(用户自定义的心跳检测消息处理方法)
	OnRemoteNotAlive OnRemoteNotAlive // User-defined method for handling remote connections that are not alive(用户自定义的远程连接不存活时的处理方法)
	HeartBeatMsgID   uint32           // User-defined ID for heartbeat detection messages(用户自定义的心跳检测消息ID)
	Router           SRouter          // User-defined business processing route for heartbeat detection messages(用户自定义的心跳检测消息业务处理路由)
}

const (
	HeartBeatDefaultMsgID uint16 = 1
)
