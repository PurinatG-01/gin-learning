package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
)

type TicketService interface {
	All() (*[]model.Ticket, error)
	Create(ticket model.Ticket) (bool, error)
}

func NewTicketService(repository repository.TicketRepository) TicketService {
	return &ticketService{repository: repository}
}

type ticketService struct {
	repository repository.TicketRepository
}

func (s *ticketService) All() (*[]model.Ticket, error) {
	tickets, err := s.repository.All()
	return tickets, err
}

func (s *ticketService) Create(ticket model.Ticket) (bool, error) {
	_, err := s.repository.Create(&ticket)
	return true, err
}
