package car

import (
	"context"
	"encoding/json"
	"log"

	"github.com/haupc/cartransplant/grpcproto"
)

func (c *carServer) FindTrip(ctx context.Context, req *grpcproto.FindTripRequest) (*grpcproto.JsonResponse, error) {
	routes, err := c.TripService.FindTrip(req.From, req.To, req.BeginLeaveTime, req.EndLeaveTime, req.Option)
	if err != nil {
		log.Printf("FindTrip - Error: %v", err)
		return nil, err
	}
	byteData, err := json.Marshal(routes)
	return &grpcproto.JsonResponse{JsonResponse: byteData}, err
}
