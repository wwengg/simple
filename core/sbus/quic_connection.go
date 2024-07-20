package sbus

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/quic-go/quic-go"
	"math/big"
	"time"
)

type QuicConnection struct {
	BaseConnection

	stream quic.Stream
}

func NewQuicConnection(cID uint64, taskHandler STaskHandler, frameDecoder SFrameDecoder, stream quic.Stream, onConnStart, onConnStop func(conn SConnection), datapack SDataPack) SConnection {
	return &QuicConnection{
		BaseConnection: BaseConnection{
			connID:         cID,
			connIdStr:      fmt.Sprintf("%d", cID),
			taskHandler:    taskHandler,
			onConnStart:    onConnStart,
			onConnStop:     onConnStop,
			datapack:       datapack,
			frameDecoder:   frameDecoder,
			sendFunc:       quicSendFunc(stream),
			readFunc:       quicReadFunc(stream),
			property:       nil,
			IOReadBuffSize: 0,
		},
		stream: stream,
	}
}

func GenerateTLSConfig() (*tls.Config, error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:     x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key),
	})
	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM := pem.EncodeToMemory(&b)

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quick"},
	}, nil
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
