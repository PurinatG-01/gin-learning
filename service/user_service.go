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
	GetPublic(id int) (model.PublicUser, error)
	GetTransactionList(userId int, page int, limit int, status model.OmiseChargeStatus, order model.OrderBy) (model.Pagination[model.TicketsTransaction], error)
	GetTicketsList(userId int, page int, limit int) (model.Pagination[model.UsersAccess], error)
	Delete(id int) (bool, error)
	Update(id int, user model.UpdateFormUser) (model.Users, error)
	IsUsernameExist(username string) (bool, error)
}

func NewUserService(userRepository repository.UserRepository, usersAccessRepository repository.UsersAccessRepository, ticketsTransactionRepository repository.TicketTransactionRepository) UserService {
	return &userService{userRepository: userRepository, usersAccessRepository: usersAccessRepository, ticketsTransactionRepository: ticketsTransactionRepository}
}

type userService struct {
	userRepository               repository.UserRepository
	usersAccessRepository        repository.UsersAccessRepository
	ticketsTransactionRepository repository.TicketTransactionRepository
}

func (s *userService) All() (*[]model.Users, error) {
	users, err := s.userRepository.All()
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
	user := s.MapFormUserToUsers(form_user, hash_pass)
	_, err := s.userRepository.Create(&user)
	return true, err
}

func (s *userService) Get(id int) (model.Users, error) {
	user, err := s.userRepository.Get(id)
	return user, err
}

func (s *userService) GetPublic(id int) (model.PublicUser, error) {
	user, err := s.userRepository.Get(id)
	public_user := s.MapUserToPublicUser(user)
	return public_user, err
}

func (s *userService) GetTicketsList(userId int, page int, limit int) (model.Pagination[model.UsersAccess], error) {
	events_pagination, err := s.usersAccessRepository.ListByUserId(userId, page, limit)
	return events_pagination, err
}

func (s *userService) GetTransactionList(userId int, page int, limit int, status model.OmiseChargeStatus, order model.OrderBy) (model.Pagination[model.TicketsTransaction], error) {
	events_pagination, err := s.ticketsTransactionRepository.ListByUserId(userId, page, limit, status, order)
	return events_pagination, err
}

func (s *userService) Delete(id int) (bool, error) {
	user := model.Users{Id: id}
	_, err := s.userRepository.Delete(&user)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *userService) Update(id int, user model.UpdateFormUser) (model.Users, error) {
	updated_user := s.MapUpdateFormUserToUsers(user)
	updated_user.Id = id
	now := time.Now()
	updated_user.UpdatedAt = &now
	_, err := s.userRepository.Update(&updated_user)
	if err != nil {
		return updated_user, err
	}
	return updated_user, nil
}

func (s *userService) IsUsernameExist(username string) (bool, error) {
	result, err := s.userRepository.IsExist("username", username)
	return result, err
}

func (s *userService) MapUserToPublicUser(user model.Users) model.PublicUser {
	public_user := model.PublicUser{
		Id:            user.Id,
		Username:      user.Username,
		DisplayName:   user.DisplayName,
		DisplayImgUrl: user.DisplayImgUrl,
		Email:         user.Email,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
	}
	return public_user
}

func (s *userService) MapFormUserToUsers(form_user model.FormUser, hash_pass []byte) model.Users {
	user := model.Users{
		Username:      form_user.Username,
		DisplayName:   form_user.DisplayName,
		DisplayImgUrl: form_user.DisplayImgUrl,
		Email:         form_user.Email,
		Password:      string(hash_pass[:]),
	}
	return user
}

func (s *userService) MapUpdateFormUserToUsers(form_user model.UpdateFormUser) model.Users {
	user := model.Users{
		DisplayName:   form_user.DisplayName,
		DisplayImgUrl: form_user.DisplayImgUrl,
		Email:         form_user.Email,
	}
	return user
}
