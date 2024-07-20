package sbus

type SFrameDecoder interface {
	Decode(buff []byte) ([][]byte, error)
}
