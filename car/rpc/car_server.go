package car

import (
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/car/service"
)

type carServer struct {
	TripService        service.TripService
	NotifyService      service.NotifyService
	DriverProvinceRepo repository.DriverProvinceRepo
}

func NewCarServer() *carServer {
	return &carServer{
		TripService:        service.GetTripService(),
		DriverProvinceRepo: repository.GetDriverProvinceRepo(),
		NotifyService:      service.GetNotifyService(),
	}
}
