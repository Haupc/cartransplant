package dto

import (
	greometry_dto "github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

type RegisterTripRequest struct {
	BeginLeaveTime int64               `json:"begin_leave_time"`
	EndLeaveTime   int64               `json:"end_leave_time"`
	From           greometry_dto.Point `json:"from"`
	To             greometry_dto.Point `json:"to"`
	MaxDistance    float32             `json:"max_distance"`
	CarID          int64               `json:"car_id"`
	FeeEachKm      int64               `json:"fee_each_km"`
	Seat           int64               `json:"seat"`
}

type TripRequest struct {
	BeginLeaveTime int64               `json:"begin_leave_time"`
	EndLeaveTime   int64               `json:"end_leave_time"`
	From           greometry_dto.Point `json:"from"`
	To             greometry_dto.Point `json:"to"`
	Opt            int32               `json:"opt"`
	Seat           int64               `json:"seat"`
	DriverTripID   int64               `json:"driver_trip_id"`
}

type FindTripResponse struct {
	ID             int64                    `json:"driver_trip_id"`
	Route          greometry_dto.RoutingDTO `json:"route"`
	Car            *grpcproto.CarObject     `json:"car"`
	UserID         string                   `json:"user_id"`
	BeginLeaveTime int64                    `json:"begin_leave_time"`
	EndLeaveTime   int64                    `json:"end_leave_time"`
	Price          int64                    `json:"price"`
	Distance       float64                  `json:"distance"`
	RemainingSeat  int32                    `json:"remaining_seat"`
}

type TripLocationInfo struct {
	From *grpcproto.Point `json:"from"`
	To   *grpcproto.Point `json:"to"`
}
