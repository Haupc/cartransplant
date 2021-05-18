package config

import (
	"log"

	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func init() {
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
}
