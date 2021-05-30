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

func Distance(from, to *grpcproto.Point) int64 {
	request := &grpcproto.RouteRequest{
		From: from,
		To:   to,
	}
	response, err := client.GetGeomClient().GetRouting(context.Background(), request)
	if err != nil {
		log.Println("Distance - Error:", err)
		return 0
	}
	var route dto.RoutingDTO
	err = json.Unmarshal(response.JsonResponse, &route)
	if err != nil {
		log.Println("Distance - Error:", err)
		return 0
	}
	return int64(route.Routes[0].Distance)
}
