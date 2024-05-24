package client

import (
	"net"

	"github.com/njayp/octopus/pkg/grpc/conn"
	"github.com/njayp/octopus/pkg/grpc/proto"
)

type myListener struct {
	srv proto.ReverseConnection_ConnectClient
}

func NewListener(srv proto.ReverseConnection_ConnectClient) *myListener {
	return &myListener{srv: srv}
}

func (l *myListener) Accept() (net.Conn, error) {
	return conn.NewReversedConn(l.srv)
}

func (l *myListener) Close() error {
	return nil
}

func (l *myListener) Addr() net.Addr {
	return &net.IPAddr{IP: net.IPv4zero}
}
