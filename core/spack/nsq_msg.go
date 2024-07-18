package spack

type SerializeType byte

const (
	// SerializeNone uses raw []byte and don't serialize/deserialize
	SerializeNone SerializeType = iota
	// JSON for payload.
	JSON
	// ProtoBuffer for payload.
	ProtoBuffer
)

// CompressType defines decompression type.
type CompressType byte

const (
	// None does not compress.
	None CompressType = iota
	// Gzip uses gzip compression.
	Gzip
	Brotli
)

// MessageType is message type of requests and responses.
type MessageType byte

const (
	// Request is message type of request
	Request MessageType = iota
	// Response is message type of response
	Response
)

type NSQMsg struct {
	//*Header
	//PkgLen        uint32  // nsq不存在粘包问题 不需要
	Cmd           uint16
	Ret           uint16
	Version       uint8
	SerializeType SerializeType
	CompressType  CompressType
	MessageType   MessageType
	Seq           uint64
	Metadata      map[string]string
	Data          []byte
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
func (m *NSQMsg) GetSerializeType() SerializeType {
	return m.SerializeType
}
func (m *NSQMsg) GetCompressType() CompressType {
	return m.CompressType
}
func (m *NSQMsg) GetMessageType() MessageType {
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
