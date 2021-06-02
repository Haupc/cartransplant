package auth

import "github.com/haupc/cartransplant/auth/service"

type authServer struct {
	AuthService service.AuthService
}

func NewAuthServer() *authServer {
	return &authServer{
		AuthService: service.GetAuthService(),
	}
}
