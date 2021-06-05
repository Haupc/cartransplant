package client

import (
	"log"

	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

var (
	conn         *grpc.ClientConn
	notifyClient grpcproto.NotifyClient
)

func GetNotifyClient() grpcproto.NotifyClient {
	if notifyClient == nil {
		var err error
		conn, err = grpc.Dial(":9003", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		notifyClient = grpcproto.NewNotifyClient(conn)
	}
	return notifyClient
}
