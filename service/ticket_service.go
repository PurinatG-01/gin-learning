package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
	"time"
)

type TicketService interface {
	All() (*[]model.Ticket, error)
	Create(ticket model.Ticket) (bool, error)
	Get(id int) (model.Ticket, error)
	Delete(id int) (bool, error)
	Update(id int, ticket model.Ticket) (bool, error)
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

func (s *ticketService) Get(id int) (model.Ticket, error) {
	ticket, err := s.repository.Get(id)
	return ticket, err
}

func (s *ticketService) Delete(id int) (bool, error) {
	ticket := model.Ticket{Id: id}
	_, err := s.repository.Delete(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Update(id int, ticket model.Ticket) (bool, error) {
	ticket.Id = id
	now := time.Now()
	ticket.UpdatedAt = &now
	_, err := s.repository.Update(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}
