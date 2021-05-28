package utils

import (
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/grpcproto"
)

func CarModelToCarRPC(car *model.Car) (carRPC *grpcproto.CarObject) {
	return &grpcproto.CarObject{
		Id:           int32(car.ID),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
	}
}

func CarRPCToCarModel(car *grpcproto.CarObject) (carModel *model.Car) {
	return &model.Car{
		ID:           int(car.Id),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
	}
}
