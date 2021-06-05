package repository

import (
	"context"
	"errors"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

// -------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- NOTIFICATION REPO ----------------------------------------------------------

type NotifyRepo interface {
	GetAllNotifyRepoByUserID(ctx context.Context, userID string, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error)
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

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- IMPLEMENT ----------------------------------------------------------

func (n *notifyRepo) GetAllNotifyRepoByUserID(ctx context.Context, userID string, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error) {
	err = n.db.WithContext(ctx).Model(&grpcproto.NotifyMessage{}).
		Where("user_id = ?", userID).
		Order("created_time DESC").
		Offset(offset).Limit(limit).
		Find(&notis).Error

	if err != nil {
		return nil, err
	}

	return notis, nil
}

func (n *notifyRepo) InsertNewNoti(ctx context.Context, notification *grpcproto.NotifyMessage) (err error) {
	// pre-exec check
	if notification == nil {
		return errors.New("Nil noti")
	}

	err = n.db.WithContext(ctx).Model(&grpcproto.NotifyMessage{}).Save(notification).Error

	if err != nil {
		return err
	}

	return nil
}
