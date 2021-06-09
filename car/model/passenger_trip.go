package model

import (
	"encoding/json"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/car/dto"
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
	Note           string `json:"note"`
}

func (p *PassengerTrip) TableName() string {
	return "passenger_trip"
}

func (p *PassengerTrip) ToGrpcListUserTripResponse(driverInfro, userInfo *grpcproto.UserProfile, car *grpcproto.CarObject) (*grpcproto.UserTrip, *dto.TripLocationInfo) {
	var locationTripInfo dto.TripLocationInfo
	json.Unmarshal([]byte(p.Location), &locationTripInfo)
	from := base.GetLocationName(locationTripInfo.From)
	to := base.GetLocationName(locationTripInfo.To)
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
		Note:           p.Note,
		// Distance:       float32(distance),
	}, &locationTripInfo
}
