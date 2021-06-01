package car

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) TakeTrip(ctx context.Context, req *grpcproto.TakeTripRequest) (*grpcproto.Bool, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	err := c.TripService.TakeTrip(md.UserID, int64(req.DriverTripID), req.BeginLeaveTime, req.EndLeaveTime, req.Seat, req.From, req.To)
	if err != nil {
		log.Printf("TakeTrip - Error: %v", err)
		return &grpcproto.Bool{Value: false}, err
	}
	return &grpcproto.Bool{Value: true}, nil
}
