package car

import (
	"context"

	"github.com/golang/glog"
	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) ListDriverTrip(ctx context.Context, req *grpcproto.ListDriverTripRequest) (*grpcproto.ListDriverTripResponse, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	glog.V(3).Infof("metadata: %v", md)
	return c.TripService.ListDriverTrip(md.UserID, req.State, req.StartDate, req.EndDate)
}
