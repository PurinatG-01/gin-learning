package repository

import (
	"errors"
	"fmt"
	model "gin-learning/models"

	"gorm.io/gorm"
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userAdapter{DB: db}
}

type userAdapter struct {
	DB *gorm.DB
}

func (s *userAdapter) All() (*[]model.User, error) {
	var users *[]model.User
	result := s.DB.Find(&users)
	return users, result.Error
}

func (s *userAdapter) Create(user *model.User) (bool, error) {
	result := s.DB.Create(user)
	return true, result.Error
}

func (s *userAdapter) Get(id int) (model.User, error) {
	var user model.User
	result := s.DB.First(&user, id)
	if result.RowsAffected == 0 {
		return user, errors.New("[GET] user id not found")
	}
	return user, result.Error
}

func (s *userAdapter) Update(user *model.User) (bool, error) {
	result := s.DB.Model(user).Updates(user)
	return true, result.Error
}

func (s *userAdapter) Delete(user *model.User) (bool, error) {
	result := s.DB.Delete(user)
	return true, result.Error
}

func (s *userAdapter) IsExist(key string, value string) (bool, error) {
	var exists bool
	err := s.DB.Model(model.User{}).
		Select("count(*) > 0").
		Where(fmt.Sprintf("%s = ?", key), value).
		Find(&exists).
		Error

	if err != nil {
		return false, err
	}
	return exists, err
}

func (s *userAdapter) GetByKey(key string, value string) (model.User, error) {
	var userStruct model.User
	result := s.DB.Where(fmt.Sprintf("%s = ?", key), value).First(&userStruct)
	fmt.Printf("\n\n> %+v, %+v \n\n", userStruct, nil)
	if result.RowsAffected != 1 {
		return userStruct, errors.New(fmt.Sprintf("%s : %s found more than 1 (rows affeceted more than 1)", key, value))
	} else if result.Error != nil {
		return userStruct, result.Error
	}
	return userStruct, nil
}
