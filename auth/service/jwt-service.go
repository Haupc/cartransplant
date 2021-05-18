package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/haupc/cartransplant/auth/dto"
	"github.com/haupc/cartransplant/cache"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var _jwtService *jwtService

// JwtService do jwt things
type JwtService interface {
	GenerateToken(userID string, roles []string) (dto.TokenDTO, error)
	ValidateToken(token string) (*jwt.Token, error)
	RefeshToken(userID string, roles []string, refeshToken string) (dto.TokenDTO, error)
}

// JwtClaim information in jwt
type jwtClaim struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey         string
	issuer            string
	tokenCacheService cache.Cache
}

// GetJwtService singleton
func GetJwtService() JwtService {
	if _jwtService == nil {
		return &jwtService{
			secretKey:         GetSecretKey(),
			issuer:            "my_issuer",
			tokenCacheService: cache.GetTokenCache(),
		}
	}

	return _jwtService
}

// GetSecretKey get secret key to use
func GetSecretKey() string {
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "my_secret"
	}
	return secret
}

func (j *jwtService) GenerateToken(userID string, roles []string) (dto.TokenDTO, error) {
	tokenDTO := dto.BuildTokenDTO()
	claims := &jwtClaim{
		userID,
		roles,
		jwt.StandardClaims{
			ExpiresAt: tokenDTO.JwtExpr,
			IssuedAt:  time.Now().Unix(),
			Issuer:    j.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	tokenDTO.JwtToken = t
	tokenDTO.RefeshToken = uuid.New().String()
	j.tokenCacheService.Set(tokenDTO.RefeshToken, userID, tokenDTO.RefeshExpr)
	return tokenDTO, err
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing method wrong %v", t.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}

func (j *jwtService) RefeshToken(userID string, roles []string, refeshToken string) (dto.TokenDTO, error) {
	checkRfToken, err := j.tokenCacheService.Get(refeshToken)
	if err != nil || checkRfToken == nil || checkRfToken != userID {
		return dto.TokenDTO{}, errors.New("Invalid refesh token")
	}
	j.tokenCacheService.Evict(refeshToken)
	return j.GenerateToken(userID, roles)
}
