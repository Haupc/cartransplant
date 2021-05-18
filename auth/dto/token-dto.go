package dto

import (
	"log"
	"os"
	"strconv"
	"time"
)

// TokenDTO ...
type TokenDTO struct {
	JwtToken    string
	RefeshToken string
	JwtExpr     int64
	RefeshExpr  time.Duration
}

func BuildTokenDTO() TokenDTO {
	jwtExpr, err := strconv.ParseInt(os.Getenv("JWT_TOKEN_EXPR"), 10, 32)
	if err != nil {
		log.Println(err)
		return TokenDTO{}
	}
	refeshExpr, err := strconv.ParseInt(os.Getenv("UUID_TOKEN_EXPR"), 10, 32)
	if err != nil {
		log.Println(err)
		return TokenDTO{}
	}
	return TokenDTO{
		JwtExpr:    time.Now().AddDate(0, 0, int(jwtExpr)).Unix(),
		RefeshExpr: time.Duration(refeshExpr) * time.Hour * 24,
	}
}
