package model

import "gorm.io/gorm"

// Trip entity
type Trip struct {
	gorm.Model
	WayJson string
	UserID  int64
}

// TableName name of table
func (r Trip) TableName() string {
	return "trip"
}
