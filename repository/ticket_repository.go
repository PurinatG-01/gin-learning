package repository

import model "gin-learning/models"

type TicketRepository interface {
	All() (*[]model.Ticket, error)
	Create(event *model.Ticket) (bool, error)
	Get(id int) (model.Ticket, error)
}
