package sface

type SMsg interface {
	GetMsgID() uint32 // Gets the ID of the message(获取消息ID)
	GetData() []byte  // Gets the content of the message(获取消息内容)
}
