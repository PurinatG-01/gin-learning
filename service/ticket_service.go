package service

import (
	"errors"
	model "gin-learning/models"
	"gin-learning/repository"
	"time"

	"github.com/kr/pretty"
)

type TicketService interface {
	All() (*[]model.Tickets, error)
	Create(ticket model.Tickets) (bool, error)
	Get(id int) (model.Tickets, error)
	Delete(id int) (bool, error)
	Update(id int, ticket model.Tickets) (bool, error)
	Purchase(form_ticket model.FormTicket, userId int) (is_success bool, err error, is_serv_err bool)
}

func NewTicketService(ticketRepository repository.TicketRepository, eventRepository repository.EventRepository, userRepository repository.UserRepository) TicketService {
	return &ticketService{ticketRepository: ticketRepository, eventRepository: eventRepository, userRepository: userRepository}
}

type ticketService struct {
	ticketRepository repository.TicketRepository
	eventRepository  repository.EventRepository
	userRepository   repository.UserRepository
}

func (s *ticketService) All() (*[]model.Tickets, error) {
	tickets, err := s.ticketRepository.All()
	return tickets, err
}

func (s *ticketService) Create(ticket model.Tickets) (bool, error) {
	_, err := s.ticketRepository.Create(&ticket)
	return true, err
}

func (s *ticketService) Get(id int) (model.Tickets, error) {
	ticket, err := s.ticketRepository.Get(id)
	return ticket, err
}

func (s *ticketService) Delete(id int) (bool, error) {
	ticket := model.Tickets{Id: id}
	_, err := s.ticketRepository.Delete(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Update(id int, ticket model.Tickets) (bool, error) {
	ticket.Id = id
	now := time.Now()
	ticket.UpdatedAt = &now
	_, err := s.ticketRepository.Update(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Purchase(form_ticket model.FormTicket, userId int) (bool, error, bool) {
	// #1 Check if eventId has available ticket left
	event, event_err := s.eventRepository.Get(form_ticket.EventId)
	if event_err != nil {
		return false, event_err, true
	}
	pretty.Print(event.AvailableTickets < form_ticket.Amount)
	// #1.1 Out of ticket => error
	if event.AvailableTickets < form_ticket.Amount {
		return false, errors.New("Out of ticket"), false
	}
	// #1.2 Success => continue
	// #2 Check if userId has enough money in wallet
	user, user_err := s.userRepository.Get(userId)
	if user_err != nil {
		return false, user_err, true
	}
	// #2.1 Not enough money => error
	if user.TotalMoney < (event.TicketPrice * form_ticket.Amount) {
		return false, errors.New("Not have enough money"), false
	}
	// #2.2 Success => continue
	// #3 Create Ticket from total amount
	// #4 Create Ticket transaction from total amount
	// #5 Create Ticket user access from total amount
	// #6 Update Event available tickets by counting from tickets

	return true, nil, false
}
