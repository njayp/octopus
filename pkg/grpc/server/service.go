package server

import (
	"context"
	"fmt"
	"net"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/conn"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"google.golang.org/grpc"
)

type Service struct {
	proto.UnimplementedReverseConnectionServer
}

func NewService() error {
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterReverseConnectionServer(s, &Service{})
	return s.Serve(lis)
}

func (s *Service) Connect(srv proto.ReverseConnection_ConnectServer) error {
	conn, err := grpc.NewClient("", grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return conn.NewReverseConn(srv)
	}))
	if err != nil {
		return err
	}
	cli := proto.NewPingerClient(conn)
	cli.Ping(srv.Context(), &proto.Empty{})
	return nil
}
