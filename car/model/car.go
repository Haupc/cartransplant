package model

import "time"

type Car struct {
	ID           int    `gorm:"Column:id; Type:int4; primarykey"`
	UserID       string `gorm:"Column:user_id; Type:varchar"`
	LicensePlate string `gorm:"Column:license_plate; Type:varchar(19)"`
	Color        string `gorm:"Column:color; Type:varchar"`
	Model        string `gorm:"Column:model; Type:varchar"`
	Deleted      bool   `gorm:"Column:deleted; Type:bool; Default:false"`
	Seat         int32  `gorm:"Column:seat"`
	VehicleBrand string `gorm:"Column:vehicle_brand"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

// TableName name of table
func (c Car) TableName() string {
	return "car"
}
