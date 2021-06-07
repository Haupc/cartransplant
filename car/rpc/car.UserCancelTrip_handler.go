package car

import (
	"context"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) CancelTrip(ctx context.Context, tripID *grpcproto.Int) (*grpcproto.Bool, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	err := c.TripService.CancelTrip(md.UserID, int32(tripID.Value))
	if err != nil {
		return nil, err
	}
	return &grpcproto.Bool{Value: true}, nil
}
