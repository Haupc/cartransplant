package car

import (
	"context"

	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) FindPendingTrip(ctx context.Context, req *grpcproto.FindPendingTripRequest) (*grpcproto.ListUserTripResponse, error) {
	return c.TripService.FindPendingTrip(req.From, req.To, req.Date, req.Seat, req.Type)
}
