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
	"gorm.io/gorm"

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
	FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, tripType, seat int32) ([]trip_dto.FindTripResponse, error)
	TakeTrip(userID string, driverTripID, beginLeaveTime, endLeaveTime int64, seat int32, from, to *grpcproto.Point) error
	ListUserTrip(userID string, state int32) (*grpcproto.ListUserTripResponse, error)
	ListDriverTrip(userID string, state, startDate, endDate int32) (*grpcproto.ListDriverTripResponse, error)
	RegisterTripUser(userID string, beginLeaveTime, endLeaveTime int64, from, to *grpcproto.Point, seat, tripType int32) error
	FindPendingTrip(seat, radius, tripType int32, rootPoint *grpcproto.Point) (*grpcproto.ListUserTripResponse, error)
	UpdateUserTrip(driverTripID, userTripID, userTripPrice int32) error
	GetPassengerTripByID(userTripID int32) (*model.PassengerTrip, error)
	CreateTripByUserTrip(tripModel *model.Trip) error
	GetLastTripID(userID string, carID, maxDistance, priceEachKm, totalSeat int32) (int32, error)
	CancelTrip(userID string, userTripID int32) error
	MarkUserTripDone(userTripID int32) error
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

func (s *tripService) MarkUserTripDone(userTripID int32) error {
	userTrip, err := s.PassengerTripRepo.FindPassengerTripByID(userTripID)
	if err != nil {
		log.Printf("MarkUserTripDone - userTrip not found - Error: %v", err)
		return err
	}
	if userTrip.State != 2 {
		log.Printf("MarkUserTripDone - Trip cant be marked as done")
		return errors.New("Trip cant be marked as done")
	}
	userTrip.State = 3
	err = s.PassengerTripRepo.Update(userTrip)
	if err != nil {
		log.Printf("MarkUserTripDone - update trip fail - Error: %v", err)
		return err
	}
	if userTrip != nil {
		remaningTrip := s.PassengerTripRepo.RemainingUserTripByTripID(int32(userTrip.TripID))
		if remaningTrip == 0 {
			return s.TripRepo.UpdateState(int32(userTrip.TripID), 2)
		}
	}
	return nil
}

func (s *tripService) CancelTrip(userID string, userTripID int32) error {
	userTripModel, err := s.PassengerTripRepo.FindPassengerTripByID(userTripID)
	if err != nil || userTripModel == nil || userTripModel.State > 2 || userTripModel.UserID != userID {
		log.Printf("CancelTrip - Trip not found - Error: %v", err)
		return err
	}
	userTripModel.State = base.TRIP_STATUS_CANCELLED
	err = s.PassengerTripRepo.Update(userTripModel)
	if err != nil {
		log.Printf("CancelTrip - update trip - Error: %v", err)
		return err
	}
	return nil
}

func (s *tripService) GetLastTripID(userID string, carID, maxDistance, priceEachKm, totalSeat int32) (int32, error) {
	tripModel := &model.Trip{
		UserID:      userID,
		CarID:       int64(carID),
		MaxDistance: maxDistance,
		FeeEachKm:   int64(priceEachKm),
		Seat:        totalSeat,
	}
	result, err := s.TripRepo.FindLastTrip(tripModel)
	if err != nil {
		return 0, err
	}
	return int32(result.ID), err
}

func (s *tripService) CreateTripByUserTrip(tripModel *model.Trip) error {
	return s.TripRepo.Create(tripModel)
}

func (s *tripService) GetPassengerTripByID(userTripID int32) (*model.PassengerTrip, error) {
	return s.PassengerTripRepo.FindPassengerTripByID(userTripID)
}

func (s *tripService) UpdateUserTrip(driverTripID, userTripID, userTripPrice int32) error {
	passengerTrip := &model.PassengerTrip{
		Model: gorm.Model{
			ID: uint(userTripID),
		},
		Price:  int64(userTripPrice),
		TripID: int64(driverTripID),
	}
	err := s.PassengerTripRepo.Update(passengerTrip)
	if err != nil {
		log.Printf("UpdateUserTrip - Error: %v", err)
		return err
	}
	return nil
}

func (s *tripService) FindPendingTrip(seat, radius, tripType int32, rootPoint *grpcproto.Point) (*grpcproto.ListUserTripResponse, error) {
	passengerTrips, err := s.PassengerTripRepo.FindPendingTrip(seat, radius, tripType, rootPoint)
	if err != nil {
		return nil, err
	}
	response := &grpcproto.ListUserTripResponse{
		UserTrip: []*grpcproto.UserTrip{},
	}
	for _, trip := range passengerTrips {
		userInfo, err := auth_client.GetAuthClient().GetUserInfo(context.Background(), &grpcproto.GetUserInfoRequest{UserID: trip.UserID})
		if err != nil {
			log.Printf("get user info - error: %v", err)
		}

		grpcUserTrip, _ := trip.ToGrpcListUserTripResponse(nil, userInfo, nil)
		response.UserTrip = append(response.UserTrip, grpcUserTrip)
	}
	return response, nil
}

func (s *tripService) RegisterTripUser(userID string, beginLeaveTime, endLeaveTime int64, from, to *grpcproto.Point, seat, tripType int32) error {
	if beginLeaveTime < time.Now().Unix() {
		beginLeaveTime = time.Now().Unix()
	}
	if beginLeaveTime > endLeaveTime {
		return errors.New("Invalid time")
	}
	locationInfo, _ := json.Marshal(trip_dto.TripLocationInfo{
		From: from,
		To:   to,
	})
	distance, _ := utils.Distance(from, to)
	passengerTrip := model.PassengerTrip{
		UserID:         userID,
		Seat:           seat,
		Location:       string(locationInfo),
		State:          1,
		BeginLeaveTime: time.Unix(beginLeaveTime, 0),
		EndLeaveTime:   time.Unix(endLeaveTime, 0),
		Price:          int64(distance) * 12,
	}
	err := s.PassengerTripRepo.Create(&passengerTrip, from)
	if err != nil {
		log.Printf("RegisterTripUser - Error: %v", err)
		return err
	}
	return nil
}

func (s *tripService) ListDriverTrip(userID string, state, startDate, endDate int32) (*grpcproto.ListDriverTripResponse, error) {
	tripModels, err := s.TripRepo.GetTripByUserID(userID, state, startDate, endDate)
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
			State:          tripModel.State,
		}
		// TODO: usertrip, state
		userTripModels, err := s.PassengerTripRepo.FindUserTrip(model.PassengerTrip{TripID: int64(tripModel.ID)})
		if err != nil {
			glog.V(3).Infof("ListDriverTrip - Error: %v", err)
			return nil, err
		}
		driverTrip.TotalUserTrip = int32(len(userTripModels))
		totalIncome := 0
		for _, userTripModel := range userTripModels {
			userInfo, err := auth_client.GetAuthClient().GetUserInfo(context.Background(), &grpcproto.GetUserInfoRequest{UserID: userTripModel.UserID})
			if err != nil {
				glog.V(3).Infof("Error getting user info: %v", err)
				return nil, err
			}
			userTripRPC, locationInfo := userTripModel.ToGrpcListUserTripResponse(nil, userInfo, nil)
			if userTripRPC != nil || locationInfo != nil {
				distance, duration := utils.Distance(locationInfo.From, locationInfo.To)
				userTripRPC.Distance = float32(distance / 1000)
				userTripRPC.Duration = float32(duration)
				driverTrip.UserTrips = append(driverTrip.UserTrips, userTripRPC)
				totalIncome += int(userTripRPC.Price)
			}
		}
		driverTrip.TotalIncome = int64(totalIncome)
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

func (s *tripService) FindTrip(from *grpcproto.Point, to *grpcproto.Point, beginLeaveTime int64, endLeaveTime int64, tripType, seat int32) ([]trip_dto.FindTripResponse, error) {
	models, err := s.TripRepo.FindTrip(from, to, tripType)
	if err != nil {
		return nil, err
	}
	// Distance calculator
	distance, _ := utils.Distance(from, to)

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
		realEndLeaveTime := float64(m.EndLeaveTime.Unix()) + subRoute.Routes[0].Duration
		remainingSeat := m.Seat - s.TripRepo.GetTakenSeatByTripID(int64(m.ID))
		if !(int64(realEndLeaveTime) < beginLeaveTime || int64(realBeginLeaveTime) > endLeaveTime) && seat < remainingSeat {
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
				Price:          int64(float64(m.FeeEachKm)*distance) / 1000,
				Distance:       distance,
				RemainingSeat:  remainingSeat,
			})
		}
	}
	return result, nil
}

func (s *tripService) TakeTrip(userID string, driverTripID, beginLeaveTime, endLeaveTime int64, seat int32, from, to *grpcproto.Point) error {
	locationInfo := trip_dto.TripLocationInfo{
		From: from,
		To:   to,
	}
	locationJson, _ := json.Marshal(locationInfo)
	distance, _ := utils.Distance(from, to)
	driverTrip, _ := s.TripRepo.GetTripByID(driverTripID)

	passengerTripModel := &model.PassengerTrip{
		UserID:         userID,
		TripID:         driverTripID,
		Seat:           seat,
		Location:       string(locationJson),
		State:          base.TRIP_STATUS_TAKEN,
		BeginLeaveTime: time.Unix(beginLeaveTime, 0),
		EndLeaveTime:   time.Unix(endLeaveTime, 0),
		Price:          int64(float64(driverTrip.FeeEachKm)*distance) / 1000,
		Type:           driverTrip.Type,
	}
	err := s.PassengerTripRepo.Create(passengerTripModel, from)
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
		var userInfo *grpcproto.UserProfile
		var carRpc *grpcproto.CarObject
		if u.State == 2 || u.State == 3 {
			driverTrip, _ := s.TripRepo.GetTripByID(u.TripID)
			userInfo, err = auth_client.GetAuthClient().GetUserInfo(context.Background(), &grpcproto.GetUserInfoRequest{UserID: u.UserID})
			if err != nil {
				log.Printf("Error getting user info: %v", err)
				return nil, err
			}
			carDB, _ := s.CarRepo.GetCarByID(context.Background(), int(driverTrip.CarID))
			carRpc = utils.CarModelToCarRPC(carDB)
		}

		userTripRPC, locationInfo := u.ToGrpcListUserTripResponse(nil, userInfo, carRpc)
		if userTripRPC != nil || locationInfo != nil {
			distance, duration := utils.Distance(locationInfo.From, locationInfo.To)
			userTripRPC.Distance = float32(distance / 1000)
			userTripRPC.Duration = float32(duration)
			response.UserTrip = append(response.UserTrip, userTripRPC)
		}
	}
	return response, nil
}
