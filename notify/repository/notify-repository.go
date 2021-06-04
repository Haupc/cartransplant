package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/grpcproto"
	"gorm.io/gorm"
)

// -------------------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- NOTIFICATION REPO ----------------------------------------------------------

type NotifyRepo interface {
	GetAllNotifyRepoByUserID(ctx context.Context, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error)
}

var (
	_notifyRepo *notifyRepo
)

// NotifyRepo interact with notify in DB
type notifyRepo struct {
	db *gorm.DB
}

func GetNotifyRepo() NotifyRepo {
	if _notifyRepo == nil || _notifyRepo.db == nil {
		_notifyRepo = &notifyRepo{
			config.GetDbConnection(),
		}
	}
	return _notifyRepo
}

// -----------------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------- IMPLEMENT ----------------------------------------------------------

func (n *notifyRepo) GetAllNotifyRepoByUserID(ctx context.Context, limit int, offset int) (notis []*grpcproto.NotifyMessage, err error) {
	userID, err := n.getUserIDFromContext(ctx)
	if err != nil {
		return nil, err
	}

	err = n.db.WithContext(ctx).Model(&grpcproto.NotifyMessage{}).
		Where("user_id LIKE ?", userID).
		Order("created_time DESC").
		Offset(offset).Limit(limit).
		Find(&notis).Error

	if err != nil {
		return nil, err
	}

	return notis, nil
}

// ------------------------------
// getUserIDFromContext get userID from context by key and check
func (n *notifyRepo) getUserIDFromContext(ctx context.Context) (userID string, err error) {
	const userIDKey = "userID"

	if ctx.Err() != nil {
		return "", ctx.Err()
	}

	userIDItf := ctx.Value(userIDKey)
	switch u := userIDItf.(type) {
	case string:
		if len(u) == 0 {
			return "", errors.New("Empty userID")
		}
		return u, nil

	default:
		return "", fmt.Errorf("userID type: %T. Value %v", userIDItf, userIDItf)
	}
}
