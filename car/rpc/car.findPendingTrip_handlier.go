package car

import (
	"context"

	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) FindPendingTrip(ctx context.Context, req *grpcproto.FindPendingTripRequest) (*grpcproto.ListUserTripResponse, error) {
	rootPoint := &grpcproto.Point{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	return c.TripService.FindPendingTrip(req.Seat, int32(req.Radius), req.Type, rootPoint)
}
