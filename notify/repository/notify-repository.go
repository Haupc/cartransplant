package repository

import (
	"github.com/haupc/cartransplant/auth/config"
	"gorm.io/gorm"
)

type NotifyRepo interface {
}

var (
	_notifyRepo *notifyRepo
)

// NotifyRepo interact with notify in DB
type notifyRepo struct {
	db *gorm.DB
}

func GetNotifyRepo() NotifyRepo {
	if _notifyRepo == nil {
		_notifyRepo = &notifyRepo{
			config.GetDbConnection(),
		}
	}
	return _notifyRepo
}
