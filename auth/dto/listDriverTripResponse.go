package dto

import "github.com/haupc/cartransplant/grpcproto"

type DriverTripResponse struct {
	UserTrips      []*grpcproto.UserTrip `json:"userTrips"`
	BeginLeaveTime int64                 `json:"beginLeaveTime"`
	EndLeaveTime   int64                 `json:"endLeaveTime"`
	From           string                `json:"from"`
	To             string                `json:"to"`
	State          int32                 `json:"state"`
	TotalSeat      int32                 `json:"totalSeat"`
	ReamaingSeat   int32                 `json:"reamaingSeat"`
	Car            *grpcproto.CarObject  `json:"car"`
	PriceEachKm    int32                 `json:"priceEachKm"`
}

func DriverTripRPCToDriverTripResponse(*grpcproto.DriverTrip) DriverTripResponse {

}
