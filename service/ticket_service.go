package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
	"time"
)

type TicketService interface {
	All() (*[]model.Tickets, error)
	Create(ticket model.Tickets) (bool, error)
	Get(id int) (model.Tickets, error)
	Delete(id int) (bool, error)
	Update(id int, ticket model.Tickets) (bool, error)
}

func NewTicketService(repository repository.TicketRepository) TicketService {
	return &ticketService{repository: repository}
}

type ticketService struct {
	repository repository.TicketRepository
}

func (s *ticketService) All() (*[]model.Tickets, error) {
	tickets, err := s.repository.All()
	return tickets, err
}

func (s *ticketService) Create(ticket model.Tickets) (bool, error) {
	_, err := s.repository.Create(&ticket)
	return true, err
}

func (s *ticketService) Get(id int) (model.Tickets, error) {
	ticket, err := s.repository.Get(id)
	return ticket, err
}

func (s *ticketService) Delete(id int) (bool, error) {
	ticket := model.Tickets{Id: id}
	_, err := s.repository.Delete(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Update(id int, ticket model.Tickets) (bool, error) {
	ticket.Id = id
	now := time.Now()
	ticket.UpdatedAt = &now
	_, err := s.repository.Update(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}
