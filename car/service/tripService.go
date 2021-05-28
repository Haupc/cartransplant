package service

import (
	"context"
	"encoding/json"
	"log"
	"time"

	trip_dto "github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/grpcproto"
)

var _tripService *tripService

type tripService struct {
	TripRepo repository.TripRepo
}

type TripService interface {
	CreateTrip(route dto.RoutingDTO, userID int32, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime int64) error
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

func (s *tripService) CreateTrip(route dto.RoutingDTO, userID int32, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime int64) error {
	timeStartTime := time.Unix(beginLeaveTime, 0)
	timeEndTime := time.Unix(endLeaveTime, 0)
	return s.TripRepo.CreateTrip(route, userID, carID, maxDistance, timeStartTime, timeEndTime)
}

func (s *tripService) FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error) {
	models, err := s.TripRepo.FindTrip(from, to)
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
		routeReq := &grpcproto.RouteRequest{
			From: route.Waypoints[0].Location.ToGrpcPoint(),
			To:   from,
		}
		response, err := client.GetGeomClient().GetRouting(context.Background(), routeReq)
		if err != nil {
			log.Println("Request err: ", err)
			return nil, err
		}
		var subRoute dto.RoutingDTO
		err = json.Unmarshal(response.JsonResponse, &subRoute)
		if err != nil {
			log.Println("Parse json err: ", err)
			return nil, err
		}
		realBeginLeaveTime := float64(m.BeginLeaveTime.Unix()) + subRoute.Routes[0].Duration
		realEndeaveTime := float64(m.EndLeaveTime.Unix()) + subRoute.Routes[0].Duration

		if !(int64(realEndeaveTime) < beginLeaveTime || int64(realBeginLeaveTime) > endLeaveTime) {
			//TODO: add car infor
			result = append(result, trip_dto.FindTripResponse{
				Route:          route,
				UserID:         m.UserID,
				BeginLeaveTime: m.BeginLeaveTime.Unix(),
				EndLeaveTime:   m.EndLeaveTime.Unix(),
			})
		}
	}
	return result, nil
}
