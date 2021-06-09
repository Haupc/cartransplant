package repository

import (
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"gorm.io/gorm"
)

type DriverProvinceRepo interface {
	Create(model *model.DriverProvince) error
	BatchCreate(models []model.DriverProvince) error
	SelectAllProvinceByDriverID(driverID string) ([]int32, error)
	BatchDelete(userID string, province []int32) error
}

var (
	_driverProvinceRepo *driverProvinceRepo
)

// DriverProvinceRepo interact with driverProvince in DB
type driverProvinceRepo struct {
	db *gorm.DB
}

func GetDriverProvinceRepo() DriverProvinceRepo {
	if _driverProvinceRepo == nil {
		_driverProvinceRepo = &driverProvinceRepo{
			config.GetDbConnection(),
		}
	}
	return _driverProvinceRepo
}

func (r *driverProvinceRepo) BatchDelete(userID string, province []int32) error {
	return r.db.Where("driver_id = ? and province_id in (?)", userID, province).Delete(model.DriverProvince{}).Error
}

func (r *driverProvinceRepo) SelectAllProvinceByDriverID(driverID string) ([]int32, error) {
	result := []int32{}
	if err := r.db.Model(&model.DriverProvince{}).Select("province_id").Where("driver_id = ?", driverID).Find(&result).Error; err != nil {
		log.Printf("SelectAllByDriverID query - Error: %v", err)
		return nil, err
	}
	return result, nil
}

func (r *driverProvinceRepo) BatchCreate(models []model.DriverProvince) error {
	return r.db.Create(&models).Error
}

func (r *driverProvinceRepo) Create(model *model.DriverProvince) error {
	return r.db.Create(model).Error
}
