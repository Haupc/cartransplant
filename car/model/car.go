package model

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	UserID       int64  `json:"user_id"`
	LicensePlate string `json:"license_plate"`
	Color        string `json:"color"`
	CarModel     string `json:"model" gorm:"column:model"`
}

// TableName name of table
func (c Car) TableName() string {
	return "car"
}
