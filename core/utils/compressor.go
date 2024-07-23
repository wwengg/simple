package utils

import (
	"github.com/google/brotli/go/cbrotli"
)

type BrotliCompressor struct{}

func (c BrotliCompressor) Zip(data []byte) ([]byte, error) {
	return cbrotli.Encode(data, cbrotli.WriterOptions{Quality: 5})
}

func (c BrotliCompressor) Unzip(data []byte) ([]byte, error) {
	return cbrotli.Decode(data)
}
