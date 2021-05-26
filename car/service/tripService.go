package service

import (
	"encoding/json"
	"log"
	"time"

	trip_dto "github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/grpcproto"
)

var _tripService *tripService

type tripService struct {
	TripRepo repository.TripRepo
}

type TripService interface {
	CreateTrip(route dto.RoutingDTO, userID int32, startTime int64) error
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error)
}

func GetTripService() TripService {
	if _tripService == nil {
		_tripService = &tripService{
			TripRepo: repository.GettripRepo(),
		}
	}
	return _tripService
}

func (s *tripService) CreateTrip(route dto.RoutingDTO, userID int32, startTime int64) error {
	timeStartTime := time.Unix(startTime, 0)
	return s.TripRepo.CreateTrip(route, userID, timeStartTime)
}

func (s *tripService) FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error) {
	models, err := s.TripRepo.FindTrip(from, to, beginLeaveTime, endLeaveTime, opt)
	if err != nil {
		return nil, err
	}
	var result []trip_dto.FindTripResponse
	for _, m := range models {
		var route dto.RoutingDTO
		err := json.Unmarshal([]byte(m.WayJson), &route)
		if err != nil {
			log.Println("Parse json err: ", err)
			return nil, err
		}
		result = append(result, trip_dto.FindTripResponse{
			Route:     route,
			UserID:    m.UserID,
			LeaveTime: m.LeaveTime,
		})
	}
	return result, nil
}
