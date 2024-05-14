package server

import (
	"context"
	"fmt"
	"net"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"google.golang.org/grpc"
)

type Service struct {
	proto.UnimplementedPingerServer
}

func NewService() error {
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	lis, err := net.Listen("tcp", url)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterPingerServer(s, &Service{})
	return s.Serve(lis)
}

func (s *Service) Ping(_ context.Context, _ *proto.Empty) (*proto.Empty, error) {
	return &proto.Empty{}, nil
}

func (s *Service) Reverse(srv proto.Pinger_ReverseServer) error {
	return nil
}
