package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/auth/service"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var auth *authController

// AuthController : do auth things
type AuthController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	RefeshToken(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	tokenCache  cache.Cache
}

// GetAuthController singleton auth
func GetAuthController() AuthController {
	if auth == nil {
		auth = &authController{
			service.GetAuthService(),
			cache.GetTokenCache(),
		}
	}
	return auth
}

func (a *authController) Register(ctx *gin.Context) {
	var registerBody dto.RegisterDTO
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &registerBody)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	success, err := a.authService.Register(registerBody.Username, registerBody.Password)
	if err != nil {
		response := utils.BuildErrorResponse("Register failed", err.Error(), registerBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse(success, "success", registerBody)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) Login(ctx *gin.Context) {
	var loginBody dto.LoginDTO
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	err := json.Unmarshal(body, &loginBody)
	if err != nil {
		respose := utils.BuildErrorResponse("Request wrong format", err.Error(), body)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, respose)
		return
	}
	token, err := auth.authService.Login(loginBody.Username, loginBody.Password)
	if err != nil {
		response := utils.BuildErrorResponse("Authen fail", err.Error(), loginBody)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	response := utils.BuildResponse(true, "success", token)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) Logout(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")
	expTimeS := int64(ctx.Value("user_info").(jwt.MapClaims)["exp"].(float64))
	remainingTime := time.Now().Unix() - expTimeS
	a.tokenCache.Set(authHeader, "-1", time.Duration(remainingTime))
	response := utils.BuildResponse(true, "success", authHeader)
	ctx.JSON(http.StatusOK, response)
}

func (a *authController) RefeshToken(ctx *gin.Context) {
	refeshToken := ctx.GetHeader("Authorization")
	userID := ctx.Value("user_info").(jwt.MapClaims)["user_id"].(string)
	token, err := a.authService.RefreshToken(userID, refeshToken)
	if err != nil {
		response := utils.BuildErrorResponse("Refresh token fail", err.Error(), nil)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}
	response := utils.BuildResponse(true, "success", token)
	ctx.JSON(http.StatusOK, response)

}
