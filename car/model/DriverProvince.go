package model

type DriverProvince struct {
	DriverID   string `json:"driver_id"`
	ProvinceID int32  `json:"province_id"`
}

func (d *DriverProvince) TableName() string {
	return "driver_province"
}
