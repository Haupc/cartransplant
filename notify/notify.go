package main

import (
	"log"
	"net"

	"github.com/haupc/cartransplant/grpcproto"
	notify "github.com/haupc/cartransplant/notify/rpc"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcproto.RegisterNotifyServer(grpcServer, notify.NewNotifyServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
