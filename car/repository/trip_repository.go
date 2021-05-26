package repository

import (
	"github.com/haupc/cartransplant/auth/config"
	"gorm.io/gorm"
)

var tripRepository *tripRepo

// tripRepo interact with db
type TripRepo interface {
}

type tripRepo struct {
	db *gorm.DB
}

// GettripRepo singleton trip repo
func GettripRepo() TripRepo {
	if tripRepository == nil {
		tripRepository = &tripRepo{
			config.GetDbConnection(),
		}
	}
	return tripRepository
}
