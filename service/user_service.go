package service

import (
	"errors"
	model "gin-learning/models"
	"gin-learning/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	All() (*[]model.Users, error)
	Create(user model.FormUser) (bool, error)
	Get(id int) (model.Users, error)
	Delete(id int) (bool, error)
	Update(id int, user model.Users) (bool, error)
	IsUsernameExist(username string) (bool, error)
}

func NewUserService(repository repository.UserRepository) UserService {
	return &userService{repository: repository}
}

type userService struct {
	repository repository.UserRepository
}

func (s *userService) All() (*[]model.Users, error) {
	users, err := s.repository.All()
	return users, err
}

func (s *userService) Create(form_user model.FormUser) (bool, error) {
	result, find_err := s.IsUsernameExist(form_user.Username)
	if find_err != nil {
		return false, find_err
	}
	if !!result {
		return false, errors.New("Username already exists")
	}
	hash_pass, crypt_err := bcrypt.GenerateFromPassword([]byte(form_user.Password), bcrypt.DefaultCost)
	if crypt_err != nil {
		return false, crypt_err

	}
	// Mapping FormUser to DB User
	user := model.Users{
		Username:      form_user.Username,
		DisplayName:   form_user.DisplayName,
		DisplayImgUrl: form_user.DisplayImgUrl,
		Email:         form_user.Email,
		Password:      string(hash_pass[:]),
	}
	_, err := s.repository.Create(&user)
	return true, err
}

func (s *userService) Get(id int) (model.Users, error) {
	user, err := s.repository.Get(id)
	return user, err
}

func (s *userService) Delete(id int) (bool, error) {
	user := model.Users{Id: id}
	_, err := s.repository.Delete(&user)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *userService) Update(id int, user model.Users) (bool, error) {
	user.Id = id
	now := time.Now()
	user.UpdatedAt = &now
	_, err := s.repository.Update(&user)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *userService) IsUsernameExist(username string) (bool, error) {
	result, err := s.repository.IsExist("username", username)
	return result, err
}
