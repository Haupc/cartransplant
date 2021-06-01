package car

import (
	"context"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) ListUserTrip(ctx context.Context, req *grpcproto.Int) (*grpcproto.ListUserTripResponse, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	state := req.GetValue()
	return c.TripService.ListUserTrip(md.UserID, int32(state))
}
