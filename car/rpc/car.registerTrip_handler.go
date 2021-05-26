package car

import (
	"context"
	"encoding/json"
	"log"

	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"
	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) RegisterTrip(ctx context.Context, req *grpcproto.RegisterTripRequest) (*grpcproto.Int, error) {
	routeReq := &grpcproto.RouteRequest{
		From: req.From,
		To:   req.To,
	}
	resp, err := client.GetGeomClient().GetRouting(ctx, routeReq)
	if err == nil {
		log.Printf("RegisterTrip - Error: %v", err)
	}
	var respObj dto.RoutingDTO
	err = json.Unmarshal(resp.JsonResponse, &respObj)
	var userID int32 = 32
	c.TripService.CreateTrip(respObj, int32(userID))
}
