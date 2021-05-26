package service

import (
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/geometry/dto"
)

var _tripService *tripService

type tripService struct {
	TripRepo repository.TripRepo
}

type TripService interface {
	CreateTrip(route dto.RoutingDTO, userID int32)
}

func GetTripService() TripService {
	if _tripService == nil {
		_tripService = &tripService{
			TripRepo: repository.GettripRepo(),
		}
	}
	return _tripService
}

func (s *tripService) CreateTrip(route dto.RoutingDTO, userID int32) {

}
