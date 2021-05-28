package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/car/model"
	"gorm.io/gorm"
)

type iCarRepo interface {
	RegisterCar(ctx context.Context, car *model.Car) (err error)
	GetAllCarByUserID(ctx context.Context, userID int) (cars []*model.Car, err error)
	GetCarByID(ctx context.Context, carID int) (car *model.Car, err error)
	UpdateCarByID(ctx context.Context, carID int, car *model.Car) (err error)
	DeleteCarByID(ctx context.Context, carID int) (err error)
}

// CarRepo interact with car in DB
type CarRepo struct {
	db *gorm.DB
}

var (
	carRepo = new(CarRepo)
	carMd   = &model.Car{}
)

func GetCarRepo() iCarRepo {
	if carRepo.db != nil {
		carRepo.db = config.GetDbConnection()
	}

	return carRepo
}

// RegisterCar ...
func (c *CarRepo) RegisterCar(ctx context.Context, car *model.Car) (err error) {
	if err = c.db.WithContext(ctx).
		Model(carMd).
		Save(car).Error; err != nil {
		return err
	}

	return nil
}

// GetAllCarByUserID ...
func (c *CarRepo) GetAllCarByUserID(ctx context.Context, userID int) (cars []*model.Car, err error) {
	rows, err := c.db.WithContext(ctx).
		Model(carMd).
		Where("user_id = ? AND deleted = ?", userID, false).Rows()
	if err != nil {
		return nil, err
	}

	cars, err = c.scanCars(ctx, c.db, rows)
	if err != nil {
		return nil, err
	}

	return cars, nil
}

func (c *CarRepo) scanCars(ctx context.Context, tx *gorm.DB, rows *sql.Rows) (cars []*model.Car, err error) {
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
func (c *CarRepo) GetCarByID(ctx context.Context, carID int) (car *model.Car, err error) {
	// alloc
	car = new(model.Car)
	if err = c.db.WithContext(ctx).
		Model(carMd).
		Where("id = ? AND deleted = ?", carID, false).
		Take(car).Error; err != nil {
		return nil, err
	}

	return car, nil
}

// UpdateCarByID ...
func (c *CarRepo) UpdateCarByID(ctx context.Context, carID int, car *model.Car) (err error) {
	if err = c.db.WithContext(ctx).
		Model(carMd).
		Where("id = ? AND deleted = ?", carID, false).
		Updates(car).Error; err != nil {
		return err
	}

	return nil
}

// DeleteCarByID ...
func (c *CarRepo) DeleteCarByID(ctx context.Context, carID int) (err error) {
	if err = c.db.WithContext(ctx).
		Model(carMd).
		Where("id = ?", carID).
		Update("deleted", true).
		Update("deleted_at", time.Now).Error; err != nil {
		return err
	}

	return nil
}
