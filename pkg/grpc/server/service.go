package server

import (
	"context"
	"fmt"
	"net"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/conn"
	"github.com/njayp/octopus/pkg/grpc/pinger"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"github.com/njayp/octopus/pkg/grpc/util"
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
	proto.RegisterPingerServer(s, &pinger.Pinger{})
	return s.Serve(lis)
}

func (s *Service) Connect(srv proto.ReverseConnection_ConnectServer) error {
	conn, err := grpc.NewClient("", util.Creds, grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		return conn.NewReversedConn(srv)
	}))
	if err != nil {
		return err
	}
	cli := proto.NewPingerClient(conn)
	_, err = cli.Ping(srv.Context(), &proto.Empty{})
	if err != nil {
		return err
	}
	println("hoo ya")
	return nil
}
