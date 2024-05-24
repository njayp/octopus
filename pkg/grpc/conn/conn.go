package conn

import (
	"bufio"
	"bytes"
	"context"
	"net"
	"time"

	"github.com/njayp/octopus/pkg/grpc/proto"
)

type stream interface {
	Send(*proto.Chunk) error
	Recv() (*proto.Chunk, error)
	Context() context.Context
}

type reversedConn struct {
	s  stream
	rw *bufio.ReadWriter
}

func NewReversedConn(srv stream) (*reversedConn, error) {
	a := make([]byte, 0)
	rw := bufio.NewReadWriter(bufio.NewReader(bytes.NewBuffer(a)), bufio.NewWriter(bytes.NewBuffer(a)))

	go func() {
		for {
			select {
			case <-srv.Context().Done():
				return
			default:
				chunk, err := srv.Recv()
				if err != nil {
					panic(err)
				}
				rw.Write(chunk.Data)
			}
		}
	}()

	return &reversedConn{s: srv, rw: rw}, nil
}

func (c *reversedConn) Read(b []byte) (int, error) {
	return c.rw.Read(b)
}

func (c *reversedConn) Write(b []byte) (int, error) {
	err := c.s.Send(&proto.Chunk{Data: b})
	if err != nil {
		return 0, err
	}
	return len(b), err
}

func (c *reversedConn) Close() error {
	return nil
}

func (c *reversedConn) LocalAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

func (c *reversedConn) RemoteAddr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}

func (c *reversedConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *reversedConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *reversedConn) SetWriteDeadline(t time.Time) error {
	return nil
}
