package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"fmt"
	"log"
	"os"
)

var db *gorm.DB

// GetDbConnection connection singleton
func GetDbConnection() *gorm.DB {
	if db == nil {
		InitConnection()
	}
	return db
}

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

// InitConnection init connection
func InitConnection() {
	var err error
	db, err = gorm.Open(postgres.Open(DbURL(BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("connect fail")
	}
}

// BuildDBConfig load config from file
func BuildDBConfig() *DBConfig {
	err := godotenv.Load("database.env")
	if err != nil {
		log.Fatalf("Error loading database.env file")
	}
	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Port:     os.Getenv("DB_PORT"),
	}
}

// DbURL build connection string
func DbURL(dbConfig *DBConfig) string {
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Ho_Chi_Minh",
		dbConfig.Host,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.DBName,
		dbConfig.Port,
	)
}
