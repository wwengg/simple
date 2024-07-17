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
)

type NSQMsg struct {
	//PkgLen        uint32  // nsq不存在粘包问题 不需要
	Cmd           uint16
	Seq           uint64
	Ret           uint16
	SerializeType byte
	CompressType  byte
	Data          []byte
}
