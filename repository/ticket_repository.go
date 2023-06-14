package repository

import model "gin-learning/models"

type TicketRepository interface {
	All() (*[]model.Tickets, error)
	Create(tickets *model.Tickets) (bool, error)
	CreateMultiple(ticketsList []model.Tickets, batchSize int) (bool, error)
	Get(id int) (model.Tickets, error)
	Update(ticket *model.Tickets) (bool, error)
	Delete(ticket *model.Tickets) (bool, error)
}
