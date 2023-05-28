package repository

import model "gin-learning/models"

type UserRepository interface {
	All() (*[]model.User, error)
	Create(event *model.User) (bool, error)
	Get(id int) (model.User, error)
	Update(ticket *model.User) (bool, error)
	Delete(ticket *model.User) (bool, error)
	IsExist(key string, value string) (bool, error)
	GetByKey(key string, value string) (model.User, error)
}
