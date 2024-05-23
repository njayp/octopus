package client

import (
	"context"
	"fmt"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReverseServer struct {
	proto.UnimplementedPingerServer
}

func (c *ReverseServer) Connect(ctx context.Context) error {
	// TODO add tls option
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		panic(err)
	}
	cli := proto.NewReverseConnectionClient(conn)
	stream, err := cli.Connect(ctx)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	proto.RegisterPingerServer(s, &ReverseServer{})
	return s.Serve(NewListener(stream))
}

func (rs *ReverseServer) Ping(ctx context.Context, _ *proto.Empty) (*proto.Empty, error) {
	println("got ping")
	return nil, nil
}
