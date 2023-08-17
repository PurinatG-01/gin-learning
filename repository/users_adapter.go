package repository

import (
	"errors"
	"fmt"
	model "gin-learning/models"
	"log"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userAdapter{DB: db}
}

type userAdapter struct {
	DB *gorm.DB
}

func (s *userAdapter) All() (*[]model.Users, error) {
	var users *[]model.Users
	result := s.DB.Find(&users)
	return users, result.Error
}

func (s *userAdapter) Create(user *model.Users) (bool, error) {
	result := s.DB.Create(user)
	return true, result.Error
}

func (s *userAdapter) Get(id int) (model.Users, error) {
	var user model.Users
	result := s.DB.First(&user, id)
	if result.RowsAffected == 0 {
		return user, errors.New("[GET] user id not found")
	}
	return user, result.Error
}

func (s *userAdapter) Update(user *model.Users) (bool, error) {
	result := s.DB.Model(user).Updates(user).Find(user)
	return true, result.Error
}

func (s *userAdapter) Delete(user *model.Users) (bool, error) {
	result := s.DB.Delete(user)
	return true, result.Error
}

func (s *userAdapter) IsExist(key string, value string) (bool, error) {
	var exists bool
	err := s.DB.Model(model.Users{}).
		Select("count(*) > 0").
		Where(fmt.Sprintf("%s = ?", key), value).
		Find(&exists).
		Error

	if err != nil {
		return false, err
	}
	return exists, err
}

func (s *userAdapter) IsAdmin(id int) (bool, error) {
	var user model.Users
	result := s.DB.First(&user, id)
	if result.RowsAffected == 0 {
		return false, errors.New("[GET] user id not found")
	} else if result.Error != nil {
		return false, result.Error
	}
	if user.IsAdmin {
		return true, nil
	}
	return false, nil
}

func (s *userAdapter) GetByKey(key string, value string) (model.Users, error) {
	var userStruct model.Users
	result := s.DB.Where(fmt.Sprintf("%s = ?", key), value).First(&userStruct)
	if result.RowsAffected != 1 {
		return userStruct, errors.New(fmt.Sprintf("%s : %s found more than 1 (rows affeceted more than 1)", key, value))
	} else if result.Error != nil {
		return userStruct, result.Error
	}
	return userStruct, nil
}

// WithTrx enables repository with transaction
func (s userAdapter) WithTrx(trxHandle *gorm.DB) UserRepository {
	if trxHandle == nil {
		log.Print("[Users] Transaction Database not found")
		return &userAdapter{DB: trxHandle}
	}
	s.DB = trxHandle
	return &userAdapter{DB: trxHandle}
}
