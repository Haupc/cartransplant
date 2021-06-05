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
	GetAllTokenByUserID(ctx context.Context, userID string) (token []*grpcproto.UserToken, err error)
	SaveUserToken(ctx context.Context, token *grpcproto.UserToken) (err error)
	GetAllNotifyRepoByUserID(ctx context.Context, userID string, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error)
	SaveNotification(ctx context.Context, notification *grpcproto.NotifyMessage) (err error)
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

func (n *notifyRepo) GetAllTokenByUserID(ctx context.Context, userID string) (token []*grpcproto.UserToken, err error) {
	err = n.db.WithContext(ctx).Model(&grpcproto.UserToken{}).
		Where("user_id = ?", userID).
		Find(&token).Error

	if err != nil {
		return nil, err
	}

	return token, nil
}
func (n *notifyRepo) SaveUserToken(ctx context.Context, token *grpcproto.UserToken) (err error) {
	// pre-exec check
	if token == nil {
		return errors.New("Nil token")
	}

	err = n.db.WithContext(ctx).Model(&grpcproto.UserToken{}).
		Save(token).Error

	if err != nil {
		return err
	}

	return nil
}

func (n *notifyRepo) GetAllNotifyRepoByUserID(ctx context.Context, userID string, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error) {
	err = n.db.WithContext(ctx).Table("notification").
		Where("user_id = ?", userID).
		Order("created_time DESC").
		Offset(offset).Limit(limit).
		Find(&notis).Error

	if err != nil {
		return nil, err
	}

	return notis, nil
}

func (n *notifyRepo) SaveNotification(ctx context.Context, notification *grpcproto.NotifyMessage) (err error) {
	// pre-exec check
	if notification == nil {
		return errors.New("Nil noti")
	}

	err = n.db.WithContext(ctx).Table("notification").Save(notification).Error

	if err != nil {
		return err
	}

	return nil
}
