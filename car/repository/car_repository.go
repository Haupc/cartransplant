package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"gorm.io/gorm"
)

type CarRepo interface {
	RegisterCar(ctx context.Context, car *model.Car) (err error)
	GetAllCarByUserID(ctx context.Context, userID string, limit int) (cars []*model.Car, err error)
	GetCarByID(ctx context.Context, carID int) (car *model.Car, err error)
	UpdateCarByID(ctx context.Context, carID int, car *model.Car) (err error)
	DeleteCarByID(ctx context.Context, carID int) (err error)
}

var (
	_carRepo *carRepo
)

// CarRepo interact with car in DB
type carRepo struct {
	db *gorm.DB
}

func GetCarRepo() CarRepo {
	if _carRepo == nil {
		_carRepo = &carRepo{
			config.GetDbConnection(),
		}
	}
	return _carRepo
}

// RegisterCar ...
func (c *carRepo) RegisterCar(ctx context.Context, car *model.Car) (err error) {
	if err = c.db.WithContext(ctx).
		Model(&model.Car{}).
		Save(car).Error; err != nil {
		return err
	}

	return nil
}

// GetAllCarByUserID ...
func (c *carRepo) GetAllCarByUserID(ctx context.Context, userID string, limit int) (cars []*model.Car, err error) {
	rows, err := c.db.WithContext(ctx).
		Model(&model.Car{}).
		Where("user_id = ? AND deleted = ? limit ?", userID, false, limit).Rows()
	if err != nil {
		return nil, err
	}

	cars, err = c.scanCars(ctx, c.db, rows)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *carRepo) scanCars(ctx context.Context, tx *gorm.DB, rows *sql.Rows) (cars []*model.Car, err error) {
	for rows.Next() {
		car := new(model.Car)
		if err = tx.WithContext(ctx).ScanRows(rows, car); err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}

	return cars, nil
}

// GetCarByID ...
func (c *carRepo) GetCarByID(ctx context.Context, carID int) (car *model.Car, err error) {
	// alloc
	car = new(model.Car)
	if err = c.db.WithContext(ctx).
		Model(&model.Car{}).
		Where("id = ? AND deleted = ?", carID, false).
		Take(car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

// UpdateCarByID ...
func (c *carRepo) UpdateCarByID(ctx context.Context, carID int, car *model.Car) (err error) {
	if err = c.db.WithContext(ctx).
		Model(&model.Car{}).
		Where("id = ? AND deleted = ?", carID, false).
		Updates(car).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCarByID ...
func (c *carRepo) DeleteCarByID(ctx context.Context, carID int) (err error) {
	if err = c.db.WithContext(ctx).
		Model(&model.Car{}).
		Where("id = ?", carID).
		Update("deleted", true).
		Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
