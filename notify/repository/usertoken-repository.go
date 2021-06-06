package repository

import (
	"github.com/golang/glog"
	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/notify/model"
	"gorm.io/gorm"
)

type UserTokenRepo interface {
	SaveUserToken(userID, token string) error
	GetAllTokenByUserID(userID string) []string
}

var (
	_userTokenRepo *userTokenRepo
)

// UserTokenRepo interact with userToken in DB
type userTokenRepo struct {
	db *gorm.DB
}

func GetUserTokenRepo() UserTokenRepo {
	if _userTokenRepo == nil {
		_userTokenRepo = &userTokenRepo{
			config.GetDbConnection(),
		}
	}
	return _userTokenRepo
}

func (n *userTokenRepo) GetAllTokenByUserID(userID string) []string {
	var result []string
	if err := n.db.Raw("select token from user_token where user_id = ?", userID).Find(&result).Error; err != nil {
		glog.V(3).Infof("GetAllTokenByUserID - error: %v", err)
		return nil
	}
	return result
}

func (n *userTokenRepo) SaveUserToken(userID, token string) error {
	userTokenModel := model.UserToken{
		UserID: userID,
		Token:  token,
	}
	return n.db.Create(&userTokenModel).Error
}
