package car

import (
	"github.com/haupc/cartransplant/car/service"
)

type carServer struct {
	TripService service.TripService
}

func NewCarServer() *carServer {
	return &carServer{
		TripService: service.GetTripService(),
	}
}
