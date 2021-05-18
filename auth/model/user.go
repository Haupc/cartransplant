package model

import (
	"time"

	"gorm.io/gorm"
)

// User Entity
type User struct {
	gorm.Model
	Username string    `json:"username"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
	DOB      time.Time `json:"dob"`
	Address  string    `json:"address"`
	Roles    []Role    `gorm:"many2many:user_role"`
	Token    string    `gorm:"-" json:"token,omitempty"`
}

// TableName name of table
func (u User) TableName() string {
	return "user"
}
