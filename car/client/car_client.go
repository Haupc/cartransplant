package client

import (
	"log"

	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

var (
	conn      *grpc.ClientConn
	carClient grpcproto.CarClient
)

func GetCarClient() grpcproto.CarClient {
	if carClient == nil {
		var err error
		conn, err = grpc.Dial(":9001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		carClient = grpcproto.NewCarClient(conn)
	}
	return carClient
}
