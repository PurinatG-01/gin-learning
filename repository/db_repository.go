package repository

import (
	model "gin-learning/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	All() (*[]model.Events, error)
	Create(event *model.Events) (bool, error)
	List(page int, limit int) (model.Pagination[model.Events], error)
	Get(int) (model.Events, error)
	Update(*model.Events) (bool, error)
	Delete(*model.Events) (bool, error)
	GetByKey(key string, value string) (model.Events, error)
	WithTrx(trxHandle *gorm.DB) EventRepository
	// Load(ctx context.Context, id string) (*model.Event, error)
	// Patch(ctx context.Context, user map[string]interface{}) (int64, error)
}

type TicketRepository interface {
	All() (*[]model.Tickets, error)
	Create(tickets *model.Tickets) (model.Tickets, error)
	CreateMultiple(ticketsList *[]model.Tickets, batchSize int) ([]model.Tickets, error)
	Get(id int) (model.Tickets, error)
	Update(ticket *model.Tickets) (bool, error)
	Delete(ticket *model.Tickets) (bool, error)
	Count(ticket *model.Tickets) (int64, error)
	WithTrx(trxHandle *gorm.DB) TicketRepository
}

type UsersAccessRepository interface {
	All() (*[]model.UsersAccess, error)
	Create(tickets *model.UsersAccess) (model.UsersAccess, error)
	CreateMultiple(users_access *[]model.UsersAccess, batchSize int) (bool, error)
	Get(id int) (model.UsersAccess, error)
	Update(ticket *model.UsersAccess) (bool, error)
	Delete(ticket *model.UsersAccess) (bool, error)
	WithTrx(trxHandle *gorm.DB) UsersAccessRepository
}

type UserRepository interface {
	All() (*[]model.Users, error)
	Create(event *model.Users) (bool, error)
	Get(id int) (model.Users, error)
	Update(ticket *model.Users) (bool, error)
	Delete(ticket *model.Users) (bool, error)
	IsExist(key string, value string) (bool, error)
	IsAdmin(id int) (bool, error)
	GetByKey(key string, value string) (model.Users, error)
	WithTrx(trxHandle *gorm.DB) UserRepository
}

type TicketTransactionRepository interface {
	// All() (*[]model.TicketsTransaction, error)
	// Get(id string) (model.TicketsTransaction, error)
	Create(ticketTransaction *model.TicketsTransaction) (model.TicketsTransaction, error)
	CreateMultiple(ticketsList *[]model.TicketsTransaction, batchSize int) ([]model.TicketsTransaction, error)
	WithTrx(trxHandle *gorm.DB) TicketTransactionRepository
}
