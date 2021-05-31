package car

import (
	"context"
	"encoding/json"
	"log"

	"github.com/haupc/cartransplant/auth/middleware"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) RegisterTrip(ctx context.Context, req *grpcproto.RegisterTripRequest) (*grpcproto.Bool, error) {
	routeReq := &grpcproto.RouteRequest{
		From: req.From,
		To:   req.To,
	}
	md := middleware.GetMetadataFromContext(ctx)
	resp, err := client.GetGeomClient().GetRouting(ctx, routeReq)
	if err != nil {
		log.Printf("RegisterTrip - Error: %v", err)
		return nil, err
	}
	var respObj dto.RoutingDTO
	err = json.Unmarshal(resp.JsonResponse, &respObj)
	if err != nil {
		log.Printf("RegisterTrip - Error: %v", err)
		return nil, err
	}

	err = c.TripService.CreateTrip(respObj, md.UserID, req.CarID, int64(req.MaxDistance), req.BeginLeaveTime, req.EndLeaveTime, req.FeeEachKm)
	if err != nil {
		log.Printf("RegisterTrip - Error: %v", err)
		return nil, err
	}
	return &grpcproto.Bool{Value: true}, nil
}
