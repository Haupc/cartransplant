package car

import (
	"context"
	"log"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) RegisterActiveZone(ctx context.Context, req *grpcproto.ActiveZone) (*grpcproto.Bool, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	registeredProvinces, err := c.DriverProvinceRepo.SelectAllProvinceByDriverID(md.UserID)
	if err != nil {
		return nil, err
	}
	// register new zone
	unregisteredProvinces := base.DiffSliceInt32(req.Provinces, registeredProvinces)
	if len(unregisteredProvinces) > 0 {
		err := c.TripService.RegisterActiveZone(md.UserID, unregisteredProvinces)
		if err != nil {
			log.Printf("RegisterActiveZone - Error: %v", err)
		}
	}
	// unregister zone
	toUnregisterProvinces := base.DiffSliceInt32(registeredProvinces, req.Provinces)
	if len(toUnregisterProvinces) > 0 {
		err := c.TripService.UnRegisterActiveZone(md.UserID, unregisteredProvinces)
		if err != nil {
			log.Printf("RegisterActiveZone - Error: %v", err)
		}
	}
	return &grpcproto.Bool{Value: true}, nil
}
