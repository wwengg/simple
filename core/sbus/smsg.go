package sbus

import (
	"github.com/wwengg/simple/core/smsg"
)

type SMsg interface {
	GetCmd() uint16 // Gets the ID of the message(获取消息ID)
	GetRet() uint16
	GetVersion() uint8
	GetSerializeType() smsg.SerializeType
	GetCompressType() smsg.CompressType
	GetMessageType() smsg.MessageType
	GetSeq() uint64
	GetMeta() map[string]string
	GetData() []byte // Gets the content of the message(获取消息内容)
}
