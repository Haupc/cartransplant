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
	TotalIncome    int64                 `json:"totalIncome"`
	TotalUserTrip  int32                 `json:"totalUserTrip"`
}

func DriverTripRPCToDriverTripResponse(driverTripRPC *grpcproto.DriverTrip) DriverTripResponse {
	return DriverTripResponse{
		UserTrips:      driverTripRPC.UserTrips,
		BeginLeaveTime: driverTripRPC.BeginLeaveTime,
		EndLeaveTime:   driverTripRPC.EndLeaveTime,
		From:           driverTripRPC.From,
		To:             driverTripRPC.To,
		State:          driverTripRPC.State,
		TotalSeat:      driverTripRPC.TotalSeat,
		ReamaingSeat:   driverTripRPC.ReamaingSeat,
		Car:            driverTripRPC.Car,
		PriceEachKm:    driverTripRPC.PriceEachKm,
		TotalIncome:    driverTripRPC.TotalIncome,
		TotalUserTrip:  driverTripRPC.TotalUserTrip,
	}
}
