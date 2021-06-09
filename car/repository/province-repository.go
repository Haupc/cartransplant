package repository

import (
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"gorm.io/gorm"
)

type ProvinceRepo interface {
	GetProvinceByID(id int32) *model.Province
	GetProvinceByIDList(id []int32) []*model.Province
}

var (
	_provinceRepo *provinceRepo
)

// ProvinceRepo interact with province in DB
type provinceRepo struct {
	db *gorm.DB
}

func GetProvinceRepo() ProvinceRepo {
	if _provinceRepo == nil {
		_provinceRepo = &provinceRepo{
			config.GetDbConnection(),
		}
	}
	return _provinceRepo
}

func (r *provinceRepo) GetProvinceByID(id int32) *model.Province {
	result := &model.Province{}
	if err := r.db.Where("id = ?", id).Find(&result).Error; err != nil {
		log.Printf("GetProvinceByID - error: %v", err)
		return nil
	}
	return result
}

func (r *provinceRepo) GetProvinceByIDList(id []int32) []*model.Province {
	result := []*model.Province{}
	if err := r.db.Where("id in (?)", id).Find(&result).Error; err != nil {
		log.Printf("GetProvinceByID - error: %v", err)
		return nil
	}
	return result
}
