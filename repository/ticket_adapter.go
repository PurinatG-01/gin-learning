package repository

import (
	"errors"
	model "gin-learning/models"

	"gorm.io/gorm"
)

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketAdapter{DB: db}
}

type ticketAdapter struct {
	DB *gorm.DB
}

func (s *ticketAdapter) All() (*[]model.Ticket, error) {
	var tickets *[]model.Ticket
	result := s.DB.Find(&tickets)
	return tickets, result.Error
}

func (s *ticketAdapter) Create(ticket *model.Ticket) (bool, error) {
	result := s.DB.Create(ticket)
	return true, result.Error
}

func (s *ticketAdapter) Get(id int) (model.Ticket, error) {
	var ticket model.Ticket
	result := s.DB.First(&ticket, id)
	if result.RowsAffected == 0 {
		return ticket, errors.New("[GET] ticket id not found")
	}
	return ticket, result.Error
}

func (s *ticketAdapter) Update(ticket *model.Ticket) (bool, error) {
	result := s.DB.Model(ticket).Updates(ticket)
	return true, result.Error
}

func (s *ticketAdapter) Delete(ticket *model.Ticket) (bool, error) {
	result := s.DB.Delete(ticket)
	return true, result.Error
}
