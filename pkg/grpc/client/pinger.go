package client

import "github.com/njayp/octopus/pkg/grpc/proto"

type Service struct {
	proto.UnimplementedPingerServer
}
