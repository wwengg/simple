package sbus

import (
	"fmt"
	"github.com/quic-go/quic-go"
	"net"
)

type QuicConnection struct {
	BaseConnection

	conn   net.PacketConn
	stream quic.Stream
}

func NewQuicConnection(IOReadBuffSize uint32, cID uint64, taskHandler STaskHandler, frameDecoder SFrameDecoder, stream quic.Stream, onConnStart, onConnStop func(conn SConnection), datapack SDataPack) SConnection {
	return &QuicConnection{
		BaseConnection: BaseConnection{
			ConnID:         cID,
			ConnIdStr:      fmt.Sprintf("%d", cID),
			TaskHandler:    taskHandler,
			OnConnStart:    onConnStart,
			OnConnStop:     onConnStop,
			Datapack:       datapack,
			FrameDecoder:   frameDecoder,
			SendFunc:       quicSendFunc(stream),
			ReadFunc:       quicReadFunc(stream),
			Property:       nil,
			IOReadBuffSize: 0,
		},
		stream: stream,
	}
}

func quicSendFunc(stream quic.Stream) func([]byte) error {
	return func(data []byte) error {
		if _, err := stream.Write(data); err != nil {
			return err
		}
		return nil
	}
}

func quicReadFunc(stream quic.Stream) func(conn SConnection, buffer []byte) (n int, err error) {
	return func(conn SConnection, buffer []byte) (n int, err error) {
		n, err = stream.Read(buffer)
		if err != nil {
			return 0, fmt.Errorf("read msg head [read datalen=%d], error = %s", n, err)
		}
		return n, nil
	}
}
