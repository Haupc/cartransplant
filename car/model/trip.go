package model

import (
	"time"

	"gorm.io/gorm"
)

// Trip entity
type Trip struct {
	gorm.Model
	WayJson        string
	UserID         int64
	BeginLeaveTime time.Time
	EndLeaveTime   time.Time
	CarID          int64
	MaxDistance    int32
	FeeEachKm      int64
	Seat           int32
}

// TableName name of table
func (r Trip) TableName() string {
	return "trip"
}
