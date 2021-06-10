package car

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) UserRegisterTrip(ctx context.Context, req *grpcproto.TakeTripRequest) (*grpcproto.Bool, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	err := c.TripService.RegisterTripUser(md.UserID, req.BeginLeaveTime, req.EndLeaveTime, req.From, req.To, req.Seat, req.Type, req.Note)
	if err != nil {
		log.Printf("UserRegisterTripHandler - Error: %v", err)
		return &grpcproto.Bool{Value: false}, err
	}
	c.NotifyService.NotifyNewTrip(req.From)
	return &grpcproto.Bool{Value: true}, nil
}
