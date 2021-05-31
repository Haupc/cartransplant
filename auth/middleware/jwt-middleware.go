package middleware

import (
	"log"
	"net/http"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/service"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Authorize apis
func Authorize() gin.HandlerFunc {
	jwtService := service.GetJwtService()
	tokenCache := cache.GetTokenCache()

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		cacheTK, err := tokenCache.Get(authHeader)
		if authHeader == "" || (cacheTK != nil && cacheTK.(string) != "") {
			response := utils.BuildErrorResponse("Authorization fail", "Token not found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			c.Set("user_info", claims)
			log.Println("user_id: ", claims["user_id"])
			log.Println("issuer: ", claims["iss"])
			log.Println("roles:", claims["roles"])
		} else {
			log.Println(err)
			response := utils.BuildErrorResponse("Invalid token", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}

func AuthorizeJWTFirebase() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utils.BuildErrorResponse("Authorization fail", "Token not found", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		log.Println("authHeader:", authHeader)
		authClient := config.GetFirebaseAuthClient()
		token, err := authClient.VerifyIDToken(c, authHeader)
		if err != nil {
			response := utils.BuildErrorResponse("Authorization fail", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		log.Println("token:", token)
		//TODO: claim
		//TODO: get extra info
	}
}
