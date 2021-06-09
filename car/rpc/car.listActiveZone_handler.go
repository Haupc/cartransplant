package car

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) ListActiveZone(ctx context.Context, req *empty.Empty) (*grpcproto.ActiveZone, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	log.Printf("ListActiveZone - metadata: %v", md)
	activeZones, err := c.TripService.ListActiveZone(md.UserID)
	if err != nil {
		log.Printf("ListActiveZone - error: %v", err)
		return nil, err
	}
	response := &grpcproto.ActiveZone{
		Provinces: activeZones,
	}
	return response, nil
}
