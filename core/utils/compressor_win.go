//go:build windows
// +build windows

package utils

import "github.com/wwengg/simple/core/slog"

type BrotliCompressor struct{}

func (c BrotliCompressor) Zip(data []byte) ([]byte, error) {
	slog.Ins().Warnf("Windows Zip")
	return data
}

func (c BrotliCompressor) Unzip(data []byte) ([]byte, error) {
	slog.Ins().Warnf("Windows Unzip")
	return data
}
