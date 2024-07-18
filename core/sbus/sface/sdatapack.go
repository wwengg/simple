package sface

type SDataPack interface {
	GetHeadLen() uint32
	Pack(msg SMsg) ([]byte, error)
	Unpack(binaryData []byte) (SMsg, error)
}
