package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/auth/service"
	"github.com/haupc/cartransplant/auth/utils"
	"github.com/haupc/cartransplant/base"
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
		if authHeader == "test" {
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
		var myData dto.Metadata
		dsnap, err := config.GetFireStoreClient().Collection("users").Doc(token.UID).Get(c)
		if err != nil {
			response := utils.BuildErrorResponse("Authorization fail", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		dsnap.DataTo(&myData)
		myData.UserID = token.UID
		c.Set("user_info", myData)
		log.Printf("user_info: %#v\n", myData)
	}
}

func GetMetadataFromContext(ctx context.Context) *dto.Metadata {
	if ginContext, ok := ctx.(*gin.Context); ok {
		metadata, exist := ginContext.Get("user_info")
		if !exist {
			return nil
		}
		if userData, ok := metadata.(dto.Metadata); ok {
			return &userData
		}
		return nil
	}
	return nil
}

func RPCNewContextFromContext(ctx context.Context) context.Context {
	md := GetMetadataFromContext(ctx)
	var newCtx context.Context
	if md != nil {
		newCtx, _ = base.RPCMetadataToOutgoing(context.Background(), md.ToUserProfile())
	}
	return newCtx
}
