package geometry

import (
	"log"
	"net"

	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	grpcproto.RegisterGeometryServer(grpcServer, newGeometryServer())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
