package service

import (
	model "gin-learning/models"
	"gin-learning/repository"

	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	LoginUser(username string, password string) (model.User, error)
}

type loginService struct {
	repository repository.UserRepository
}

func NewLoginService(repository repository.UserRepository) LoginService {
	return &loginService{
		repository: repository,
	}
}

func (s *loginService) LoginUser(username string, password string) (model.User, error) {
	user, get_err := s.repository.GetByKey("username", username)
	if get_err != nil {
		return model.User{}, get_err
	}
	hashPassword := []byte(user.Password)
	// Comparing the password with the hash
	if compare_err := bcrypt.CompareHashAndPassword(hashPassword, []byte(password)); compare_err != nil {
		return model.User{}, compare_err
	}

	return user, nil
}
