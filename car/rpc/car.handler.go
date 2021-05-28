package car

import (
	"context"

	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/grpcproto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// RegisterCar ...
func (c *carServer) RegisterCar(ctx context.Context, car *grpcproto.CarObject) (b *grpcproto.Bool, err error) {
	var (
		fixedUserID = 32
		carDB       = &model.Car{
			UserID:       fixedUserID,
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
func (c *carServer) ListMyCar(ctx context.Context, userID *grpcproto.Int) (resp *grpcproto.ListCarResponse, err error) {
	carsDB, err := repository.GetCarRepo().GetAllCarByUserID(ctx, int(userID.Value))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// create resp
	respCars := make([]*grpcproto.CarObject, 0, len(carsDB))
	for i := range carsDB {
		respCars = append(respCars, c.carModelToCarRPC(carsDB[i]))
	}

	return &grpcproto.ListCarResponse{Cars: respCars}, nil
}

// UpdateCar ...
func (c *carServer) UpdateCar(ctx context.Context, car *grpcproto.CarObject) (b *grpcproto.Bool, err error) {
	if err = repository.GetCarRepo().UpdateCarByID(ctx, int(car.Id), c.carRPCToCarModel(car)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.Bool{Value: true}, nil
}

// DeleteCar ...
func (c *carServer) DeleteCar(ctx context.Context, carID *grpcproto.Int) (b *grpcproto.Bool, err error) {
	if err = repository.GetCarRepo().DeleteCarByID(ctx, int(carID.Value)); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &grpcproto.Bool{Value: true}, nil
}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- UTILITIES ----------------------------------------------------------

func (c *carServer) carModelToCarRPC(car *model.Car) (carRPC *grpcproto.CarObject) {
	return &grpcproto.CarObject{
		Id:           int32(car.ID),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
	}
}

func (c *carServer) carRPCToCarModel(car *grpcproto.CarObject) (carModel *model.Car) {
	return &model.Car{
		ID:           int(car.Id),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
	}
}
