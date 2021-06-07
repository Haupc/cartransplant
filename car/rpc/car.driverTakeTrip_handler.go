package car

import (
	"context"
	"encoding/json"
	"log"

	"github.com/haupc/cartransplant/base"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) DriverTakeTrip(ctx context.Context, req *grpcproto.DriverTakeTripRequest) (*grpcproto.Bool, error) {
	md := base.RPCMetadataFromIncoming(ctx)
	// them trip cho driver
	userTrip, err := c.TripService.GetPassengerTripByID(req.UserTripID)
	if err != nil {
		return nil, err
	}
	var routeReq grpcproto.RouteRequest
	err = json.Unmarshal([]byte(userTrip.Location), &routeReq)
	if err != nil {
		log.Printf("DriverTakeTrip - Error: %v", err)
		return nil, err
	}
	resp, err := client.GetGeomClient().GetRouting(ctx, &routeReq)
	var respObj dto.RoutingDTO
	err = json.Unmarshal(resp.JsonResponse, &respObj)
	if err != nil {
		log.Printf("DriverTakeTrip - Error: %v", err)
		return nil, err
	}
	totalSeat := req.RemainingSeat + userTrip.Seat
	err = c.TripService.CreateTrip(respObj, md.UserID, int64(req.CarID), int64(req.MaxDistance), userTrip.BeginLeaveTime, userTrip.EndLeaveTime, int64(req.PriceEachKm), totalSeat)
	if err != nil {
		log.Printf("RegisterTrip - Error: %v", err)
		return nil, err
	}
	// sua gia tien, state cho user
	tripID, err := c.TripService.GetLastTripID(md.UserID, req.CarID, req.MaxDistance, req.PriceEachKm, totalSeat)
	if err != nil {
		return nil, err
	}
	err = c.TripService.UpdateUserTrip(tripID, req.UserTripID, req.UserTripPrice)
	if err != nil {
		return nil, err
	}
	return &grpcproto.Bool{Value: true}, nil
}
