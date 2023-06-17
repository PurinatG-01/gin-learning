package repository

import model "gin-learning/models"

type EventRepository interface {
	All() (*[]model.Events, error)
	Create(event *model.Events) (bool, error)
	Get(int) (model.Events, error)
	Update(*model.Events) (bool, error)
	Delete(*model.Events) (bool, error)
	GetByKey(key string, value string) (model.Events, error)
	// Load(ctx context.Context, id string) (*model.Event, error)
	// Patch(ctx context.Context, user map[string]interface{}) (int64, error)
}

type TicketRepository interface {
	All() (*[]model.Tickets, error)
	Create(tickets *model.Tickets) (bool, error)
	CreateMultiple(ticketsList []model.Tickets, batchSize int) (bool, error)
	Get(id int) (model.Tickets, error)
	Update(ticket *model.Tickets) (bool, error)
	Delete(ticket *model.Tickets) (bool, error)
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
}

type TicketTransactionRepository interface {
	// All() (*[]model.TicketsTransaction, error)
	// Get(id string) (model.TicketsTransaction, error)
	Create(ticketTransaction *model.TicketsTransaction) (bool, error)
}
