package dto

import greometry_dto "github.com/haupc/cartransplant/geometry/dto"

type RegisterTripRequest struct {
	StartTime int64               `json:"start_time"`
	From      greometry_dto.Point `json:"from"`
	To        greometry_dto.Point `json:"to"`
}
