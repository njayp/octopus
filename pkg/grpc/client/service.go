package client

import (
	"context"
	"fmt"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	client proto.ReverseConnectionClient
	server proto.PingerServer
}

func NewClient() *Client {
	// TODO add tls option
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	conn, err := grpc.NewClient(url, opts...)
	if err != nil {
		panic(err)
	}
	return &Client{client: proto.NewReverseConnectionClient(conn)}
}

func (c *Client) Connect(ctx context.Context) error {
	stream, err := c.client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	st
}
