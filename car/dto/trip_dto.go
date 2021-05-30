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
	MaxDistance    int32               `json:"max_distance"`
	CarID          int64               `json:"car_id"`
	FeeEachKm      int64               `json:"fee_each_km"`
}

type FindTripRequest struct {
	BeginLeaveTime int64               `json:"begin_leave_time"`
	EndLeaveTime   int64               `json:"end_leave_time"`
	From           greometry_dto.Point `json:"from"`
	To             greometry_dto.Point `json:"to"`
	Opt            int32               `json:"opt"`
}

type FindTripResponse struct {
	Route          greometry_dto.RoutingDTO `json:"route"`
	Car            *grpcproto.CarObject     `json:"car"`
	UserID         int64                    `json:"user_id"`
	BeginLeaveTime int64                    `json:"begin_leave_time"`
	EndLeaveTime   int64                    `json:"end_leave_time"`
	Price          int64                    `json:"price"`
}
