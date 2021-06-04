package service

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/golang/glog"
	auth_client "github.com/haupc/cartransplant/auth/client"
	"github.com/haupc/cartransplant/base"
	trip_dto "github.com/haupc/cartransplant/car/dto"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/car/repository"
	"github.com/haupc/cartransplant/car/utils"
	"github.com/haupc/cartransplant/geometry/client"
	"github.com/haupc/cartransplant/geometry/dto"

	"github.com/haupc/cartransplant/grpcproto"
)

var _tripService *tripService

type tripService struct {
	TripRepo          repository.TripRepo
	CarRepo           repository.CarRepo
	PassengerTripRepo repository.PassengerTripRepo
}

type TripService interface {
	CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime, priceEachKm int64, seat int32) error
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, opt int32) ([]trip_dto.FindTripResponse, error)
	TakeTrip(userID string, driverTripID, beginLeaveTime, endLeaveTime int64, seat int32, from, to *grpcproto.Point) error
	ListUserTrip(userID string, state int32) (*grpcproto.ListUserTripResponse, error)
	ListDriverTrip(userID string, limit int32) (*grpcproto.ListDriverTripResponse, error)
}

func GetTripService() TripService {
	if _tripService == nil {
		_tripService = &tripService{
			TripRepo:          repository.GettripRepo(),
			CarRepo:           repository.GetCarRepo(),
			PassengerTripRepo: repository.GetPassengerTripRepo(),
		}
	}
	return _tripService
}

func (s *tripService) ListDriverTrip(userID string, limit int32) (*grpcproto.ListDriverTripResponse, error) {
	tripModels, err := s.TripRepo.GetTripByUserID(userID, limit)
	if err != nil {
		return nil, err
	}
	response := &grpcproto.ListDriverTripResponse{
		Trips: []*grpcproto.DriverTrip{},
	}
	for _, tripModel := range tripModels {
		var route dto.RoutingDTO
		json.Unmarshal([]byte(tripModel.WayJson), &route)
		takenSeat := s.TripRepo.GetTakenSeatByTripID(int64(tripModel.ID))
		carModel, err := s.CarRepo.GetCarByID(context.Background(), int(tripModel.CarID))
		if err != nil {
			glog.V(3).Infof("ListDriverTrip - Error: %v", err)
			return nil, err
		}

		driverTrip := &grpcproto.DriverTrip{
			UserTrips:      []*grpcproto.UserTrip{},
			BeginLeaveTime: tripModel.BeginLeaveTime.Unix(),
			EndLeaveTime:   tripModel.EndLeaveTime.Unix(),
			From:           route.Waypoints[0].Name,
			To:             route.Waypoints[1].Name,
			TotalSeat:      tripModel.Seat,
			ReamaingSeat:   tripModel.Seat - takenSeat,
			PriceEachKm:    int32(tripModel.FeeEachKm),
			Car:            utils.CarModelToCarRPC(carModel),
		}
		// TODO: usertrip, state
		userTripModels, err := s.PassengerTripRepo.FindUserTrip(model.PassengerTrip{TripID: int64(tripModel.ID)})
		if err != nil {
			glog.V(3).Infof("ListDriverTrip - Error: %v", err)
			return nil, err
		}
		for _, userTripModel := range userTripModels {
			userInfo, err := auth_client.GetAuthClient().GetUserInfo(context.Background(), &grpcproto.GetUserInfoRequest{UserID: u.UserID})
			if err != nil {
				glog.V(3).Infof("Error getting user info: %v", err)
				return nil, err
			}
			userTripRPC := userTripModel.ToGrpcListUserTripResponse(userInfo, nil)
			if userTripRPC != nil {
				driverTrip.UserTrips = append(driverTrip.UserTrips, userTripRPC)
			}
		}
		response.Trips = append(response.Trips, driverTrip)
	}
	return response, nil
}

func (s *tripService) CreateTrip(route dto.RoutingDTO, userID string, carID int64, maxDistance int64, beginLeaveTime, endLeaveTime, priceEachKm int64, seat int32) error {
	if beginLeaveTime < time.Now().Unix() {
		beginLeaveTime = time.Now().Unix()
	}
	if beginLeaveTime > endLeaveTime {
		return errors.New("Invalid time")
	}
	timeStartTime := time.Unix(beginLeaveTime, 0)
	timeEndTime := time.Unix(endLeaveTime, 0)
	carModel, err := s.CarRepo.GetCarByID(context.Background(), int(carID))
	if err != nil {
		return err
	}
	if carModel == nil {
		return errors.New("Car not existed")

	}
	return s.TripRepo.CreateTrip(route, userID, carID, maxDistance, timeStartTime, timeEndTime, priceEachKm, seat)
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
				ID:             int64(m.ID),
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

func (s *tripService) TakeTrip(userID string, driverTripID, beginLeaveTime, endLeaveTime int64, seat int32, from, to *grpcproto.Point) error {
	locationInfo, _ := json.Marshal(trip_dto.TripLocationInfo{
		From: from,
		To:   to,
	})
	distance := utils.Distance(from, to)
	driverTrip, _ := s.TripRepo.GetTripByID(driverTripID)

	passengerTripModel := &model.PassengerTrip{
		UserID:         userID,
		TripID:         driverTripID,
		Seat:           seat,
		Location:       string(locationInfo),
		State:          base.TRIP_STATUS_TAKEN,
		BeginLeaveTime: time.Unix(beginLeaveTime, 0),
		EndLeaveTime:   time.Unix(endLeaveTime, 0),
		Price:          distance * driverTrip.FeeEachKm / 1000,
	}
	err := s.PassengerTripRepo.Create(passengerTripModel)
	return err
}

func (s *tripService) ListUserTrip(userID string, state int32) (*grpcproto.ListUserTripResponse, error) {
	var userTrips []model.PassengerTrip
	var err error
	if state == -1 {
		userTrips, err = s.PassengerTripRepo.FindHistoryTrip(userID)
		if err != nil {
			return nil, err
		}
	} else {

		passengerTripModel := model.PassengerTrip{
			UserID: userID,
			State:  state,
		}
		userTrips, err = s.PassengerTripRepo.FindUserTrip(passengerTripModel)
		if err != nil {
			return nil, err
		}
	}
	response := &grpcproto.ListUserTripResponse{
		UserTrip: []*grpcproto.UserTrip{},
	}
	for _, u := range userTrips {
		driverTrip, _ := s.TripRepo.GetTripByID(u.TripID)
		userInfo, err := auth_client.GetAuthClient().GetUserInfo(context.Background(), &grpcproto.GetUserInfoRequest{UserID: u.UserID})
		if err != nil {
			log.Printf("Error getting user info: %v", err)
			return nil, err
		}
		carDB, _ := s.CarRepo.GetCarByID(context.Background(), int(driverTrip.CarID))
		carRpc := utils.CarModelToCarRPC(carDB)

		userTripRPC := u.ToGrpcListUserTripResponse(userInfo, carRpc)
		if userTripRPC != nil {
			response.UserTrip = append(response.UserTrip)
		}
	}
	return response, nil
}
