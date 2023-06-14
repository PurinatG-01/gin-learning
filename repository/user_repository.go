package repository

import model "gin-learning/models"

type UserRepository interface {
	All() (*[]model.Users, error)
	Create(event *model.Users) (bool, error)
	Get(id int) (model.Users, error)
	Update(ticket *model.Users) (bool, error)
	Delete(ticket *model.Users) (bool, error)
	IsExist(key string, value string) (bool, error)
	IsAdmin(id int) (bool, error)
	GetByKey(key string, value string) (model.Users, error)
}
