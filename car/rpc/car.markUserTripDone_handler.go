package car

import (
	"context"

	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) MarkUserTripDone(ctx context.Context, userTripID *grpcproto.Int) (*grpcproto.Bool, error) {
	err := c.TripService.MarkUserTripDone(int32(userTripID.Value))
	if err != nil {
		return nil, err
	}
	return &grpcproto.Bool{Value: true}, nil
}
