package repository

import (
	"github.com/golang/glog"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/notify/model"
	"gorm.io/gorm"
)

type NotifyRepo interface {
	SaveNotification(noti *model.Notification) error
	GetAllNotifyByUserID(userID string) ([]*model.Notification, error)
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
func (n *notifyRepo) SaveNotification(noti *model.Notification) error {
	return n.db.Create(noti).Error
}

func (n *notifyRepo) GetAllNotifyByUserID(userID string) ([]*model.Notification, error) {
	var result []*model.Notification
	if err := n.db.Where("user_id = ?", userID).Find(&result).Error; err != nil {
		glog.V(3).Infof("GetAllNotifyByUserID - Error: %v", err)
		return nil, err
	}
	return result, nil
}
