package model

import "gorm.io/gorm"

// Role entity
type Role struct {
	gorm.Model
	Name        string
	Permissions []Permission `gorm:"many2many:role_permission"`
}

// TableName name of table
func (r Role) TableName() string {
	return "role"
}
