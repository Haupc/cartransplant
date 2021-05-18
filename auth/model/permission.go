package model

import "gorm.io/gorm"

// Permission Entity
type Permission struct {
	gorm.Model
	Name  string
	Roles []Role `gorm:"many2many:role_permission"`
}

// TableName name of table
func (p Permission) TableName() string {
	return "permission"
}
