package server

import (
	"bufio"
	"bytes"
	"net"
	"time"

	"github.com/njayp/octopus/pkg/grpc/proto"
)

type myConn struct {
	srv proto.ReverseConnection_ConnectServer
	rw  bufio.ReadWriter
}

func newReverseConn(srv proto.ReverseConnection_ConnectServer) (*myConn, error) {
	a := make([]byte, 0)
	rw := bufio.NewReadWriter(bufio.NewReader(bytes.NewBuffer(a)), bufio.NewWriter(bytes.NewBuffer(a)))

	go func() {
		chunk, _ := srv.Recv()
		rw.Write(chunk.Data)
	}()

	return &myConn{}, nil
}

func (c *myConn) Read(b []byte) (int, error) {
	return c.rw.Read(b)
}

func (c *myConn) Write(b []byte) (int, error) {
	err := c.srv.Send(&proto.Chunk{Data: b})
	if err != nil {
		return 0, err
	}
	return len(b), err
}

func (c *myConn) Close() error {
	return nil
}

func (c *myConn) LocalAddr() net.Addr {
	return nil
}

func (c *myConn) RemoteAddr() net.Addr {
	return nil
}

func (c *myConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *myConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *myConn) SetWriteDeadline(t time.Time) error {
	return nil
}
