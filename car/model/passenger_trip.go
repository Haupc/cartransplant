package model

import (
	"encoding/json"
	"time"

	"github.com/haupc/cartransplant/car/dto"
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
}

func (p *PassengerTrip) TableName() string {
	return "passenger_trip"
}

func (p *PassengerTrip) ToGrpcListUserTripResponse() *grpcproto.UserTrip {
	var locationTripInfo dto.TripLocationInfo
	json.Unmarshal([]byte(p.Location), &locationTripInfo)
	return &grpcproto.UserTrip{
		Id:             int32(p.ID),
		BeginLeaveTime: p.BeginLeaveTime.Unix(),
		EndLeaveTime:   p.EndLeaveTime.Unix(),
		From:           locationTripInfo.From,
		To:             locationTripInfo.To,
		State:          p.State,
	}
}
