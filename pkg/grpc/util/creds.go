package util

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var Creds = grpc.WithTransportCredentials(insecure.NewCredentials())
