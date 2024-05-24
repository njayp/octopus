package pinger

import (
	"context"

	"github.com/njayp/octopus/pkg/grpc/proto"
)

type Pinger struct {
	proto.UnimplementedPingerServer
}

func (p *Pinger) Ping(ctx context.Context, _ *proto.Empty) (*proto.Empty, error) {
	println("got ping")
	return nil, nil
}
