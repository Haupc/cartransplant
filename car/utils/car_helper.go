package utils

import (
	"context"
	"encoding/json"
	"log"

	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func CarModelToCarRPC(car *model.Car) (carRPC *grpcproto.CarObject) {
	return &grpcproto.CarObject{
		Id:           int32(car.ID),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
		VehicleBrand: car.VehicleBrand,
		Seat:         car.Seat,
	}
}

func CarRPCToCarModel(car *grpcproto.CarObject) (carModel *model.Car) {
	return &model.Car{
		ID:           int(car.Id),
		LicensePlate: car.LicensePlate,
		Color:        car.Color,
		Model:        car.Model,
		Seat:         car.Seat,
		VehicleBrand: car.VehicleBrand,
	}
}

func Distance(from, to *grpcproto.Point) (float64, float64) {
	request := &grpcproto.RouteRequest{
		From: from,
		To:   to,
	}
	response, err := client.GetGeomClient().GetRouting(context.Background(), request)
	if err != nil {
		log.Println("Distance - Error:", err)
		return 0, 0
	}
	var route dto.RoutingDTO
	err = json.Unmarshal(response.JsonResponse, &route)
	if err != nil {
		log.Println("Distance - Error:", err)
		return 0, 0
	}
	return route.Routes[0].Distance, route.Routes[0].Duration
}
