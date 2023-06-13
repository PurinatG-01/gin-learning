package repository

import model "gin-learning/models"

type TicketRepository interface {
	All() (*[]model.Tickets, error)
	Create(event *model.Tickets) (bool, error)
	Get(id int) (model.Tickets, error)
	Update(ticket *model.Tickets) (bool, error)
	Delete(ticket *model.Tickets) (bool, error)
}
