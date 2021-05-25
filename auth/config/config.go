package config

import "github.com/haupc/cartransplant/base"

// SECRET_KEY=my_secret
// JWT_TOKEN_EXPR=24
// UUID_TOKEN_EXPR=7

type authConfig struct {
	UUIDTokenExpr  int32
	JwtTokenExpr   int32
	PostgresConfig *base.PostgresConfig
}
