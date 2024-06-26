package client

import (
	"context"
	"fmt"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/pinger"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"github.com/njayp/octopus/pkg/grpc/util"
	"google.golang.org/grpc"
)

type ReverseServer struct {
	pinger.Pinger
}

func (c *ReverseServer) Connect(ctx context.Context) error {
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	conn, err := grpc.NewClient(url, util.Creds)
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
