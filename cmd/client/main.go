package main

import (
	"context"
	"fmt"

	"github.com/njayp/octopus/pkg/config"
	"github.com/njayp/octopus/pkg/grpc/client"
	"github.com/njayp/octopus/pkg/grpc/proto"
	"github.com/njayp/octopus/pkg/grpc/util"
	"google.golang.org/grpc"
)

func main() {
	reverse()
}

func ping() {
	url := fmt.Sprintf("%s:%v", config.Env.Address, config.Env.Port)
	conn, err := grpc.NewClient(url, util.Creds)
	if err != nil {
		panic(err)
	}
	cli := proto.NewPingerClient(conn)
	_, err = cli.Ping(context.Background(), &proto.Empty{})
	if err != nil {
		panic(err)
	}
	println("got ping")
}

func reverse() {
	cli := client.ReverseServer{}
	err := cli.Connect(context.Background())
	if err != nil {
		println(err)
	}
}
