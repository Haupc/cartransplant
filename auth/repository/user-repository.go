package repository

import (
	"errors"
	"log"

	"github.com/haupc/cartransplant/auth/config"
	"github.com/haupc/cartransplant/auth/model"

	"gorm.io/gorm"
)

var userRepository *userRepo

// UserRepo interact with db
type UserRepo interface {
	CreateUser(user *model.User) (bool, error)
	FindByUserModel(user *model.User) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	FindByUserAndPassword(username, password string) (*model.User, error)
	FindByID(id int) (*model.User, error)
}

type userRepo struct {
	db *gorm.DB
}

// GetUserRepo singleton user repo
func GetUserRepo() UserRepo {
	if userRepository == nil {
		userRepository = &userRepo{
			config.GetDbConnection(),
		}
	}
	return userRepository
}

func (ur *userRepo) CreateUser(user *model.User) (bool, error) {
	if err := userRepository.db.Create(user).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (ur *userRepo) FindByUserModel(userModel *model.User) (*model.User, error) {
	var user model.User
	if err := userRepository.db.Where(userModel).Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) FindByUsername(username string) (*model.User, error) {
	var user model.User
	var err error
	if err = userRepository.db.Raw("select * from public.user where username = ?", username).Take(&user).Error; err != nil {
		err = errors.New("wrong username or password")
	}

	return &user, err
}

func (ur *userRepo) FindByUserAndPassword(username, password string) (*model.User, error) {
	var user model.User
	var err error
	if err = userRepository.db.Raw("select * from public.user where username = ? and password = ?", username, password).Take(&user).Error; err != nil {
		err = errors.New("wrong username or password")
	}

	return &user, err
}

// FindByID find user by id
func (ur *userRepo) FindByID(id int) (*model.User, error) {
	var user model.User
	if err := userRepository.db.Where("id = ?", id).Preload("Roles").Find(&user).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &user, nil
}
