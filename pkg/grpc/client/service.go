package client

import (
	"github.com/njayp/octopus/pkg/grpc/proto"
)

func Ping() {
	// make new pingerClient here
	client := proto.NewPingerClient(conn)
}
