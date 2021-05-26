package dto

import (
	"time"

	greometry_dto "github.com/haupc/cartransplant/geometry/dto"
)

type RegisterTripRequest struct {
	StartTime int64               `json:"start_time"`
	From      greometry_dto.Point `json:"from"`
	To        greometry_dto.Point `json:"to"`
}

type FindTripRequest struct {
	BeginLeaveTime int64               `json:"begin_leave_time"`
	EndLeaveTime   int64               `json:"end_leave_time"`
	From           greometry_dto.Point `json:"from"`
	To             greometry_dto.Point `json:"to"`
	Opt            int32               `json:"opt"`
}

type FindTripResponse struct {
	Route     greometry_dto.RoutingDTO `json:"route"`
	UserID    int64                    `json:"user_id"`
	LeaveTime time.Time                `json:"leave_time"`
}
