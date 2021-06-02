package model

import (
	"context"
	"encoding/json"
	"log"
	"time"

	geomdto "github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

type PassengerTrip struct {
	gorm.Model
	UserID         string    `json:"user_id"`
	TripID         int64     `json:"trip_id"`
	Seat           int32     `json:"seat"`
	Location       string    `json:"location"`
	State          int32     `json:"state"`
	BeginLeaveTime time.Time `json:"begin_leave_time"`
	EndLeaveTime   time.Time `json:"end_leave_time"`
	Price          int64     `json:"price"`
}

func (p *PassengerTrip) TableName() string {
	return "passenger_trip"
}

func (p *PassengerTrip) ToGrpcListUserTripResponse(userInfo *grpcproto.UserProfile, car *grpcproto.CarObject) *grpcproto.UserTrip {
	var locationTripInfo dto.TripLocationInfo
	json.Unmarshal([]byte(p.Location), &locationTripInfo)
	addressFrom, err := client.GetGeomClient().GetCurrentAddress(context.Background(), locationTripInfo.From)
	if err != nil {
		log.Printf("ToGrpcListUserTripResponse - GetCurrentAddress error: %v", err)
		return &grpcproto.UserTrip{}
	}
	var addressParsed geomdto.SearchAddressResponse
	json.Unmarshal(addressFrom.JsonResponse, &addressParsed)

	from := addressParsed.DisplayName

	addressTo, err := client.GetGeomClient().GetCurrentAddress(context.Background(), locationTripInfo.From)
	if err != nil {
		log.Printf("ToGrpcListUserTripResponse - GetCurrentAddress error: %v", err)
		return &grpcproto.UserTrip{}
	}
	json.Unmarshal(addressTo.JsonResponse, &addressParsed)
	to := addressParsed.DisplayName
	return &grpcproto.UserTrip{
		Id:             int32(p.ID),
		BeginLeaveTime: p.BeginLeaveTime.Unix(),
		EndLeaveTime:   p.EndLeaveTime.Unix(),
		From:           from,
		To:             to,
		State:          p.State,
		Driver:         userInfo,
		Car:            car,
		Price:          p.Price,
	}
}
