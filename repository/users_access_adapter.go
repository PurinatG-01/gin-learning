package repository

import (
	"errors"
	model "gin-learning/models"
	"log"

	"gorm.io/gorm"
)

func NewUsersAccessRepository(db *gorm.DB) UsersAccessRepository {
	return &usersAccessAdapter{DB: db}
}

type usersAccessAdapter struct {
	DB *gorm.DB
}

func (s *usersAccessAdapter) All() (*[]model.UsersAccess, error) {
	var users_access *[]model.UsersAccess
	result := s.DB.Find(&users_access)
	return users_access, result.Error
}

func (s *usersAccessAdapter) Create(users_access *model.UsersAccess) (model.UsersAccess, error) {
	result := s.DB.Create(users_access)
	return *users_access, result.Error
}

func (s *usersAccessAdapter) CreateMultiple(users_access_list *[]model.UsersAccess, batchSize int) (bool, error) {
	result := s.DB.CreateInBatches(users_access_list, batchSize)
	return true, result.Error
}

func (s *usersAccessAdapter) Get(id int) (model.UsersAccess, error) {
	var users_access model.UsersAccess
	result := s.DB.First(&users_access, id)
	if result.RowsAffected == 0 {
		return users_access, errors.New("[GET] users_access id not found")
	}
	return users_access, result.Error
}

func (s *usersAccessAdapter) Update(users_access *model.UsersAccess) (bool, error) {
	result := s.DB.Model(users_access).Updates(users_access)
	return true, result.Error
}

func (s *usersAccessAdapter) Delete(users_access *model.UsersAccess) (bool, error) {
	result := s.DB.Delete(users_access)
	return true, result.Error
}

// WithTrx enables repository with transaction
func (s usersAccessAdapter) WithTrx(trxHandle *gorm.DB) UsersAccessRepository {
	if trxHandle == nil {
		log.Print("[UsersAccess] Transaction Database not found")
		return &usersAccessAdapter{DB: trxHandle}
	}
	s.DB = trxHandle
	return &usersAccessAdapter{DB: trxHandle}
}
