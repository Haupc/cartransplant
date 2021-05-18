package service

import (
	"strconv"

	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/auth/repository"
	"github.com/haupc/cartransplant/cache"
)

var auth *authService

// AuthService interface
type AuthService interface {
	Login(username, password string) (interface{}, error)
	RefreshToken(userID, refreshToken string) (interface{}, error)
}
type authService struct {
	userRepo   repository.UserRepo
	jwtService JwtService
	userCache  cache.Cache
}

func (a *authService) Login(username, password string) (interface{}, error) {
	user, err := auth.userRepo.FindByUserAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	userInfo, err := a.userCache.Get(user.ID)
	var roles []string
	if uInfo, ok := userInfo.(dto.UserDTO); ok {
		roles = uInfo.Roles
	}
	token, err := auth.jwtService.GenerateToken(strconv.Itoa(int(user.ID)), roles)
	return token, err
}

func (a *authService) RefreshToken(userID, refreshToken string) (interface{}, error) {
	userInfo, err := a.userCache.Get(userID)
	if err != nil {
		return nil, err
	}

	return auth.jwtService.RefeshToken(userID, userInfo.(dto.UserDTO).Roles, refreshToken)
}

// GetAuthService authservice singleton
func GetAuthService() AuthService {
	if auth == nil {
		auth = &authService{
			repository.GetUserRepo(),
			GetJwtService(),
			cache.GetUserCache(),
		}
	}
	return auth
}
