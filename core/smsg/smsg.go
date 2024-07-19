package smsg

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
