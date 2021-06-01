package car

import (
	"context"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/car/utils"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterCar ...
func (c *carServer) RegisterCar(ctx context.Context, car *grpcproto.CarObject) (b *grpcproto.Bool, err error) {
	md := base.RPCMetadataFromIncoming(ctx)
	var (
		carDB = &model.Car{
			UserID:       md.UserID,
			LicensePlate: car.LicensePlate,
			Color:        car.Color,
			Model:        car.Model,
		}
	)

	if err := repository.GetCarRepo().RegisterCar(ctx, carDB); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.Bool{Value: true}, nil
}

// ListMyCar ...
func (c *carServer) ListMyCar(ctx context.Context, limit *grpcproto.Int) (resp *grpcproto.ListCarResponse, err error) {
	md := base.RPCMetadataFromIncoming(ctx)
	carsDB, err := repository.GetCarRepo().GetAllCarByUserID(ctx, md.UserID, int(limit.Value))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// create resp
	respCars := make([]*grpcproto.CarObject, 0, len(carsDB))
	for i := range carsDB {
		respCars = append(respCars, utils.CarModelToCarRPC(carsDB[i]))
	}

	return &grpcproto.ListCarResponse{Cars: respCars}, nil
}

// UpdateCar ...
func (c *carServer) UpdateCar(ctx context.Context, car *grpcproto.CarObject) (b *grpcproto.Bool, err error) {
	if err = repository.GetCarRepo().UpdateCarByID(ctx, int(car.Id), utils.CarRPCToCarModel(car)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.Bool{Value: true}, nil
}

// DeleteCar ...
func (c *carServer) DeleteCar(ctx context.Context, req *grpcproto.DeleteCarRequest) (b *grpcproto.Bool, err error) {
	if err = repository.GetCarRepo().DeleteCarByID(ctx, req.Ids); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.Bool{Value: true}, nil
}
