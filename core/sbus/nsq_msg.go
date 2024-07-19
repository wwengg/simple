package sbus

import (
	"github.com/wwengg/simple/core/smsg"
)

type NSQMsg struct {
	//*Header
	//PkgLen        uint32  // nsq不存在粘包问题 不需要
	Cmd           uint16
	Ret           uint16
	Version       uint8
	SerializeType smsg.SerializeType
	CompressType  smsg.CompressType
	MessageType   smsg.MessageType
	Seq           uint64
	Metadata      map[string]string
	Data          []byte
}

func NewNSQMsg(Cmd uint16, ret uint16, sType smsg.SerializeType, md map[string]string, data []byte) *NSQMsg {
	return &NSQMsg{
		Cmd:           Cmd,
		Ret:           ret,
		Version:       1,
		SerializeType: sType,
		CompressType:  smsg.Gzip,
		MessageType:   smsg.Response,
		Seq:           123456789812322,
		Metadata:      md,
		Data:          data,
	}
}

func (m *NSQMsg) GetCmd() uint16 {
	return m.Cmd
}
func (m *NSQMsg) GetRet() uint16 {
	return m.Ret
}
func (m *NSQMsg) GetVersion() uint8 {
	return m.Version
}
func (m *NSQMsg) GetSerializeType() smsg.SerializeType {
	return m.SerializeType
}
func (m *NSQMsg) GetCompressType() smsg.CompressType {
	return m.CompressType
}
func (m *NSQMsg) GetMessageType() smsg.MessageType {
	return m.MessageType
}

func (m *NSQMsg) GetSeq() uint64 {
	return m.Seq
}

func (m *NSQMsg) GetMeta() map[string]string {
	return m.Metadata
}

func (m *NSQMsg) GetData() []byte {
	return m.Data
}
