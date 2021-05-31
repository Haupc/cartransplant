package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	trip_dto "github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/car/utils"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/grpcproto"
)

var _tripService *tripService

type tripService struct {
	TripRepo repository.TripRepo
	CarRepo  repository.CarRepo
}

type TripService interface {
	CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime, priceEachKm int64) error
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error)
}

func GetTripService() TripService {
	if _tripService == nil {
		_tripService = &tripService{
			TripRepo: repository.GettripRepo(),
			CarRepo:  repository.GetCarRepo(),
		}
	}
	return _tripService
}

func (s *tripService) CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime, priceEachKm int64) error {
	timeStartTime := time.Unix(beginLeaveTime, 0)
	timeEndTime := time.Unix(endLeaveTime, 0)
	carModel, err := s.CarRepo.GetCarByID(context.Background(), int(carID))
	if err != nil {
		return err
	}
	if carModel == nil {
		return errors.New("Car not existed")

	}
	return s.TripRepo.CreateTrip(route, userID, carID, maxDistance, timeStartTime, timeEndTime, priceEachKm)
}

func (s *tripService) FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error) {
	models, err := s.TripRepo.FindTrip(from, to)
	if err != nil {
		return nil, err
	}
	// Distance calculator
	distance := utils.Distance(from, to)

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
			carModel, err := s.CarRepo.GetCarByID(context.Background(), int(m.CarID))
			if err != nil {
				log.Println("Parse json err: ", err)
				return nil, err
			}
			result = append(result, trip_dto.FindTripResponse{
				Route:          route,
				UserID:         m.UserID,
				Car:            utils.CarModelToCarRPC(carModel),
				BeginLeaveTime: m.BeginLeaveTime.Unix(),
				EndLeaveTime:   m.EndLeaveTime.Unix(),
				Price:          m.FeeEachKm * distance / 1000,
			})
		}
	}
	return result, nil
}
