package repository

import (
	"fmt"
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

type PassengerTripRepo interface {
	Create(model *model.PassengerTrip, startPoint *grpcproto.Point) error
	FindUserTrip(model model.PassengerTrip) ([]model.PassengerTrip, error)
	FindHistoryTrip(userID string) ([]model.PassengerTrip, error)
	FindPendingTrip(seat, radius, tripType int32, rootPoint *grpcproto.Point) ([]model.PassengerTrip, error)
	Update(model *model.PassengerTrip) error
	FindPassengerTripByID(userTripID int32) (*model.PassengerTrip, error)
	RemainingUserTripByTripID(tripID int32) int32
}

var (
	_passengerTripRepo *passengerTripRepo
)

// PassengerTripRepo interact with passengerTrip in DB
type passengerTripRepo struct {
	db *gorm.DB
}

func GetPassengerTripRepo() PassengerTripRepo {
	if _passengerTripRepo == nil {
		_passengerTripRepo = &passengerTripRepo{
			config.GetDbConnection(),
		}
	}
	return _passengerTripRepo
}

func (r *passengerTripRepo) RemainingUserTripByTripID(tripID int32) int32 {
	var result int32
	if err := r.db.Raw("select count(*) from passenger_trip where trip_id = ? and state = 2", tripID).Scan(&result).Error; err != nil {
		log.Printf("count query - Error: %v", err)
		return 0
	}
	return result
}

func (r *passengerTripRepo) FindPassengerTripByID(userTripID int32) (*model.PassengerTrip, error) {
	var result *model.PassengerTrip
	if err := r.db.Where("id = ?", userTripID).Find(&result).Error; err != nil {
		log.Printf("FindPassengerTripByID - Error: %v", err)
		return nil, err
	}
	return result, nil
}

func (r *passengerTripRepo) Update(passengerTrip *model.PassengerTrip) error {
	return r.db.Save(passengerTrip).Error
}

func (r *passengerTripRepo) FindPendingTrip(seat, radius, tripType int32, rootPoint *grpcproto.Point) ([]model.PassengerTrip, error) {
	postgisPoint := makePoint(rootPoint)
	query := fmt.Sprintf("select * from passenger_trip where st_distance(%s, start_point) <= ? and seat <= ? and type = ? and state = 1", postgisPoint)
	var result []model.PassengerTrip
	if err := r.db.Raw(query, radius, seat, tripType).Find(&result).Error; err != nil {
		log.Printf("FindPendingTrip - Error: %v", err)
	}
	return result, nil
}

func (r *passengerTripRepo) Create(userTripModel *model.PassengerTrip, startPoint *grpcproto.Point) error {
	start_point := makePoint(startPoint)
	query := fmt.Sprintf("insert into passenger_trip (user_id, trip_id, seat, location, state, begin_leave_time, end_leave_time, price, start_point, type) values (?, ?, ?, ?, ?, ?, ?, ?, %s, ?)", start_point)
	if err := r.db.Exec(query, userTripModel.UserID, userTripModel.TripID, userTripModel.Seat, userTripModel.Location, userTripModel.State, userTripModel.BeginLeaveTime, userTripModel.EndLeaveTime, userTripModel.Price, userTripModel.Type).Error; err != nil {
		return err
	}
	return nil
}

func (r *passengerTripRepo) FindUserTrip(userTripModel model.PassengerTrip) ([]model.PassengerTrip, error) {
	var result []model.PassengerTrip
	if err := r.db.Find(&result, userTripModel).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (r *passengerTripRepo) FindHistoryTrip(userID string) ([]model.PassengerTrip, error) {
	var result []model.PassengerTrip
	if err := r.db.Where("user_id = ? and state >= 3", userID).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil

}
