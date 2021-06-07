package model

import (
	"context"
	"encoding/json"
	"log"

	geomdto "github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

type PassengerTrip struct {
	gorm.Model
	UserID         string `json:"user_id"`
	TripID         int64  `json:"trip_id"`
	Seat           int32  `json:"seat"`
	Location       string `json:"location"`
	State          int32  `json:"state"`
	BeginLeaveTime int64  `json:"begin_leave_time"`
	EndLeaveTime   int64  `json:"end_leave_time"`
	Price          int64  `json:"price"`
	Type           int32  `json:"type"`
}

func (p *PassengerTrip) TableName() string {
	return "passenger_trip"
}

func (p *PassengerTrip) ToGrpcListUserTripResponse(driverInfro, userInfo *grpcproto.UserProfile, car *grpcproto.CarObject) (*grpcproto.UserTrip, *dto.TripLocationInfo) {
	var locationTripInfo dto.TripLocationInfo
	json.Unmarshal([]byte(p.Location), &locationTripInfo)
	addressFrom, err := client.GetGeomClient().GetCurrentAddress(context.Background(), locationTripInfo.From)
	if err != nil {
		log.Printf("ToGrpcListUserTripResponse - GetCurrentAddress error: %v", err)
		return nil, nil
	}
	var addressParsed geomdto.SearchAddressResponse
	json.Unmarshal(addressFrom.JsonResponse, &addressParsed)

	from := addressParsed.DisplayName

	addressTo, err := client.GetGeomClient().GetCurrentAddress(context.Background(), locationTripInfo.To)
	if err != nil {
		log.Printf("ToGrpcListUserTripResponse - GetCurrentAddress error: %v", err)
		return nil, nil
	}
	json.Unmarshal(addressTo.JsonResponse, &addressParsed)
	to := addressParsed.DisplayName
	return &grpcproto.UserTrip{
		Id:             int32(p.ID),
		BeginLeaveTime: p.BeginLeaveTime,
		EndLeaveTime:   p.EndLeaveTime,
		From:           from,
		To:             to,
		State:          p.State,
		Driver:         driverInfro,
		User:           userInfo,
		Car:            car,
		Price:          p.Price,
		Type:           p.Type,
		Seat:           p.Seat,
		// Distance:       float32(distance),
	}, &locationTripInfo
}
