package model

import (
	"gorm.io/gorm"
)

// Trip entity
type Trip struct {
	gorm.Model
	WayJson        string
	UserID         string
	BeginLeaveTime int64
	EndLeaveTime   int64
	CarID          int64
	MaxDistance    float32
	FeeEachKm      int64
	Seat           int32
	State          int32
	Type           int32
}

// TableName name of table
func (r Trip) TableName() string {
	return "trip"
}
