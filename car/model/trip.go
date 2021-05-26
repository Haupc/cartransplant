package model

import (
	"time"

	"gorm.io/gorm"
)

// Trip entity
type Trip struct {
	gorm.Model
	WayJson   string
	UserID    int64
	LeaveTime time.Time
}

// TableName name of table
func (r Trip) TableName() string {
	return "trip"
}
