package main

import "github.com/njayp/octopus/pkg/grpc/server"

func main() {
	err := server.NewService()
	if err != nil {
		println(err)
	}
}
