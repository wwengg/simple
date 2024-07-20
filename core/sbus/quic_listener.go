package sbus

import (
	"context"
	"crypto/tls"
	"github.com/quic-go/quic-go"
	"net"
	"time"
)

type QuicConn struct {
	conn  net.PacketConn
	qconn quic.Connection

	stream quic.Stream
}

// Read implements the Conn Read method.
func (c *QuicConn) Read(b []byte) (int, error) {
	return c.stream.Read(b)
}

// Write implements the Conn Write method.
func (c *QuicConn) Write(b []byte) (int, error) {
	return c.stream.Write(b)
}

// LocalAddr returns the local network address.
func (c *QuicConn) LocalAddr() net.Addr {
	return c.qconn.LocalAddr()
}

// RemoteAddr returns the remote network address.
func (c *QuicConn) RemoteAddr() net.Addr {
	return c.qconn.RemoteAddr()
}

// Close closes the connection.
func (c *QuicConn) Close() error {
	if c.stream != nil {
		return c.stream.Close()
	}

	return nil
}

// SetDeadline sets the deadline associated with the listener. A zero time value disables the deadline.
func (c *QuicConn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline implements the Conn SetReadDeadline method.
func (c *QuicConn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline implements the Conn SetWriteDeadline method.
func (c *QuicConn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}

type QuicListener struct {
	conn       net.PacketConn
	quicServer *quic.Listener
}

func NewQuicListener(c net.PacketConn, tlsConf *tls.Config, quicConfig *quic.Config) (*QuicListener, error) {
	ln, err := quic.Listen(c, tlsConf, quicConfig)
	if err != nil {
		return nil, err
	}

	return &QuicListener{
		conn:       c,
		quicServer: ln,
	}, nil
}

func (q *QuicListener) Accept() (net.Conn, error) {
	return q.AcceptContext(context.Background())
}

// AcceptContext waits for and returns the next connection to the listener.
func (q *QuicListener) AcceptContext(ctx context.Context) (net.Conn, error) {
	conn, err := q.quicServer.Accept(ctx)
	if err != nil {
		return nil, err
	}
	stream, err := conn.AcceptStream(ctx)
	if err != nil {
		return nil, err
	}

	qconn := &QuicConn{
		conn:   q.conn,
		qconn:  conn,
		stream: stream,
	}

	return qconn, nil
}

// Close closes the listener.
// Any blocked Accept operations will be unblocked and return errors.
func (q *QuicListener) Close() error {
	return q.quicServer.Close()
}

// Addr returns the listener's network address.
func (q *QuicListener) Addr() net.Addr {
	return q.quicServer.Addr()
}
