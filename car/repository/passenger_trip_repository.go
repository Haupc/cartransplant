package repository

import (
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"gorm.io/gorm"
)

type PassengerTripRepo interface {
	Create(model *model.PassengerTrip) error
	FindUserTrip(model model.PassengerTrip) ([]model.PassengerTrip, error)
	FindHistoryTrip(userID string) ([]model.PassengerTrip, error)
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

func (r *passengerTripRepo) Create(userTripModel *model.PassengerTrip) error {
	if err := r.db.Create(userTripModel).Error; err != nil {
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
