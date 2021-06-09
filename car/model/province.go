package model

type Province struct {
	ID          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Topic       string `json:"topic"`
}
