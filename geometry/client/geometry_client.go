package client

import (
	"log"

	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

var (
	conn       *grpc.ClientConn
	geomClient grpcproto.GeometryClient
)

func GetGeomClient() grpcproto.GeometryClient {
	if geomClient != nil {
		var err error
		conn, err = grpc.Dial(":9000", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		geomClient = grpcproto.NewGeometryClient(conn)
	}
	return geomClient
}
