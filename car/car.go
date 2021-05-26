package main

import (
	"log"
	"net"

	car "github.com/haupc/cartransplant/car/rpc"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcproto.RegisterCarServer(grpcServer, car.NewCarServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
