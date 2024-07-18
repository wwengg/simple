package sface

import "github.com/wwengg/simple/core/spack"

type SMsg interface {
	GetCmd() uint16 // Gets the ID of the message(获取消息ID)
	GetRet() uint16
	GetVersion() uint8
	GetSerializeType() spack.SerializeType
	GetCompressType() spack.CompressType
	GetMessageType() spack.MessageType
	GetSeq() uint64
	GetMeta() map[string]string
	GetData() []byte // Gets the content of the message(获取消息内容)
}
