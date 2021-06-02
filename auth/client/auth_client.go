package client

import (
	"log"

	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	authClient grpcproto.AuthClient
)

func GetAuthClient() grpcproto.AuthClient {
	if authClient == nil {
		var err error
		conn, err = grpc.Dial(":9002", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		authClient = grpcproto.NewAuthClient(conn)
	}
	return authClient
}
