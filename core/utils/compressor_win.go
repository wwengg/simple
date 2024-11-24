//go:build windows
// +build windows

package utils

type BrotliCompressor struct{}

func (c BrotliCompressor) Zip(data []byte) ([]byte, error) {
	return data, nil
}

func (c BrotliCompressor) Unzip(data []byte) ([]byte, error) {
	return data, nil
}
