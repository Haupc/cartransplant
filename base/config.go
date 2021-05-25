package base

// DB_HOST=localhost
// DB_PORT=5432
// DB_USER=postgres
// DB_PASSWORD=123456
// DB_NAME=auth
type PostgresConfig struct {
	UserName string
	Host     string
	Port     int32
	Password string
	DbName   string
}
