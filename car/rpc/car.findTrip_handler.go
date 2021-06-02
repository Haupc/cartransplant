package car

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) FindTrip(ctx context.Context, req *grpcproto.FindTripRequest) (*grpcproto.JsonResponse, error) {
	routes, err := c.TripService.FindTrip(req.From, req.To, req.BeginLeaveTime, req.EndLeaveTime, req.Option)
	if err != nil {
		log.Printf("FindTrip - Error: %v", err)
		return nil, err
	}
	if routes == nil || len(routes) == 0 {
		return nil, errors.New("No trip found")
	}
	byteData, err := json.Marshal(routes)
	return &grpcproto.JsonResponse{JsonResponse: byteData}, err
}
