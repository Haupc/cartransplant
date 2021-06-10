package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

type PassengerTripRepo interface {
	Create(model *model.PassengerTrip, startPoint, endPoint *grpcproto.Point) error
	FindUserTrip(model model.PassengerTrip) ([]model.PassengerTrip, error)
	FindHistoryTrip(userID string) ([]model.PassengerTrip, error)
	FindPendingTrip(from, to, date, seat int32, tripType []int32) ([]model.PassengerTrip, error)
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

func (r *passengerTripRepo) FindPendingTrip(from, to, date, seat int32, tripType []int32) ([]model.PassengerTrip, error) {
	var result []model.PassengerTrip
	query := r.db.Model(&model.PassengerTrip{}).Where("state = ?", 1)

	now := time.Now().Unix()
	today := time.Unix(now-now%86400, 0).Add(-7 * time.Hour).Unix()

	if int64(date) < today {
		query.Where("begin_leave_time >=  ?", today)
	} else {
		startDate := date
		endDate := time.Unix(int64(date), 0).AddDate(0, 0, 1).Unix()
		if date >= int32(today) && date < int32(today)+86400 {
			startDate = int32(today)
			endDate = today + 86400
		}
		query.Where("begin_leave_time between ? and ? or end_leave_time between ? and ?", startDate, endDate, startDate, endDate)

	}
	if seat > 0 {
		query = query.Where("seat <= ?", seat)
	}

	if len(tripType) > 0 {
		query = query.Where("type in (?)", tripType)
	}

	if from > 0 {
		subquery := r.db.Model(&model.Province{}).Select("way").Where("id = ?", from)
		query = query.Where("st_contains((?) , start_point)", subquery)
	}

	if to > 0 {
		subquery := r.db.Model(&model.Province{}).Select("way").Where("id = ?", to)
		query = query.Where("st_contains((?) , end_point)", subquery)
	}

	if err := query.Find(&result).Error; err != nil {
		log.Printf("FindPassengerTripByID query - Error: %v", err)
		return nil, err
	}

	return result, nil
}

func (r *passengerTripRepo) Create(userTripModel *model.PassengerTrip, startPoint, endPoint *grpcproto.Point) error {
	start_point := makePoint(startPoint)
	end_point := makePoint(endPoint)
	query := fmt.Sprintf("insert into passenger_trip (user_id, trip_id, seat, location, state, begin_leave_time, end_leave_time, price, start_point, end_point, type, note) values (?, ?, ?, ?, ?, ?, ?, ?, %s, %s, ?, ?)", start_point, end_point)
	if err := r.db.Exec(query, userTripModel.UserID, userTripModel.TripID, userTripModel.Seat, userTripModel.Location, userTripModel.State, userTripModel.BeginLeaveTime, userTripModel.EndLeaveTime, userTripModel.Price, userTripModel.Type, userTripModel.Note).Error; err != nil {
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
