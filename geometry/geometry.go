package main

import (
	"log"
	"net"

	geometry "github.com/haupc/cartransplant/geometry/rpc"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcproto.RegisterGeometryServer(grpcServer, geometry.NewGeometryServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
