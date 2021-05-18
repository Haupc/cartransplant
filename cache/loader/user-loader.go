package loader

import (
	"encoding/json"
	"strconv"

	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/auth/repository"
)

var _userLoader *userLoader

type userLoader struct {
	userRepo repository.UserRepo
}

func (u *userLoader) Load(key string) (string, error) {
	var userInfo dto.UserDTO
	id, err := strconv.Atoi(key)
	if err != nil {
		return "", err
	}
	user, err := u.userRepo.FindByID(id)
	userInfo.Name = user.Name
	userInfo.Roles = func() []string {
		roles := []string{}
		for _, r := range user.Roles {
			roles = append(roles, r.Name)
		}
		return roles
	}()
	jsonReturn, err := json.Marshal(userInfo)

	return string(jsonReturn), nil
}

// GetUserCacheLoader ...
func GetUserCacheLoader() CacheLoader {
	if _userLoader == nil {
		_userLoader = &userLoader{
			repository.GetUserRepo(),
		}
	}
	return _userLoader
}
