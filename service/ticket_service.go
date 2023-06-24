package service

import (
	"errors"
	model "gin-learning/models"
	"gin-learning/repository"
	"time"

	"gorm.io/gorm"
)

type TicketService interface {
	All() (*[]model.Tickets, error)
	Create(ticket model.Tickets) (bool, error)
	Get(id int) (model.Tickets, error)
	Delete(id string) (bool, error)
	Update(id string, ticket model.Tickets) (bool, error)
	Purchase(form_ticket model.FormTicket, userId int) (ticket model.Tickets, err error, is_serv_err bool)
	MapFormTicketToTickets(form_ticket model.FormTicket, event model.Events, userId int) model.Tickets
	WithTrx(trxHandle *gorm.DB) TicketService
}

func NewTicketService(ticketRepository repository.TicketRepository, eventRepository repository.EventRepository, userRepository repository.UserRepository, ticketTransactionRepository repository.TicketTransactionRepository, usersAccessRepository repository.UsersAccessRepository) TicketService {
	return &ticketService{ticketRepository: ticketRepository, eventRepository: eventRepository, userRepository: userRepository, ticketTransactionRepository: ticketTransactionRepository, usersAccessRepository: usersAccessRepository}
}

type ticketService struct {
	ticketRepository            repository.TicketRepository
	eventRepository             repository.EventRepository
	userRepository              repository.UserRepository
	ticketTransactionRepository repository.TicketTransactionRepository
	usersAccessRepository       repository.UsersAccessRepository
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

func (s *ticketService) Delete(id string) (bool, error) {
	ticket := model.Tickets{Id: id}
	_, err := s.ticketRepository.Delete(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Update(id string, ticket model.Tickets) (bool, error) {
	ticket.Id = id
	now := time.Now()
	ticket.UpdatedAt = &now
	_, err := s.ticketRepository.Update(&ticket)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *ticketService) Purchase(form_ticket model.FormTicket, userId int) (model.Tickets, error, bool) {
	// #1 Check if eventId has available ticket left
	event, event_err := s.eventRepository.Get(form_ticket.EventId)
	if event_err != nil {
		return model.Tickets{}, event_err, true
	}
	// #1.1 Out of ticket => error
	if *event.AvailableTickets < form_ticket.Amount {
		return model.Tickets{}, errors.New("Out of ticket"), false
	}
	// #1.2 Success => continue
	// #2 Check if userId has enough money in wallet
	user, user_err := s.userRepository.Get(userId)
	if user_err != nil {
		return model.Tickets{}, user_err, true
	}
	// #2.1 Not enough money => error
	if user.TotalMoney < (event.TicketPrice * form_ticket.Amount) {
		return model.Tickets{}, errors.New("Not have enough money"), false
	}
	// #2.2 Success => continue
	// #3 Create Ticket from total amount
	ticket := s.MapFormTicketToTickets(form_ticket, event, userId)
	ticket, ticket_err := s.ticketRepository.Create(&ticket)
	if ticket_err != nil {
		return ticket, ticket_err, true
	}
	// #4 Create Ticket transaction from total amount
	ticket_transaction := model.TicketsTransaction{TicketId: ticket.Id, PurchaserId: userId, EventId: event.Id}
	_, ticket_transaction_err := s.ticketTransactionRepository.Create(&ticket_transaction)
	if ticket_transaction_err != nil {
		return ticket, ticket_transaction_err, true
	}
	// #5 Create Ticket user access from total amount
	users_access := model.UsersAccess{TicketId: ticket.Id, UserId: userId, EventId: event.Id}
	_, users_access_err := s.usersAccessRepository.Create(&users_access)
	if users_access_err != nil {
		return ticket, users_access_err, true
	}
	// #6 Update Event available tickets by counting from tickets
	ticket_count, ticket_count_err := s.ticketRepository.Count(&model.Tickets{EventId: event.Id})
	if ticket_count_err != nil {
		return ticket, ticket_count_err, true
	}
	available_tickets := int(int64(event.TotalTickets) - ticket_count)
	event.AvailableTickets = &available_tickets
	_, update_event_err := s.eventRepository.Update(&event)
	if update_event_err != nil {
		return ticket, update_event_err, true
	}
	user.TotalMoney = user.TotalMoney - (event.TicketPrice * form_ticket.Amount)
	_, update_user_err := s.userRepository.Update(&user)
	if update_user_err != nil {
		return ticket, update_user_err, true
	}
	return ticket, nil, false
}

func (s *ticketService) MapFormTicketToTickets(form_ticket model.FormTicket, event model.Events, userId int) model.Tickets {
	return model.Tickets{
		EventId: event.Id,
		OwnerId: userId,
	}
}

func (s *ticketService) WithTrx(trxHandle *gorm.DB) TicketService {
	trxTicketRepository := s.ticketRepository.WithTrx(trxHandle)
	trxEventRepository := s.eventRepository.WithTrx(trxHandle)
	trxUserRepository := s.userRepository.WithTrx(trxHandle)
	trxTicketTransactionRepository := s.ticketTransactionRepository.WithTrx(trxHandle)
	trxUsersAccessRepository := s.usersAccessRepository.WithTrx(trxHandle)
	return &ticketService{ticketRepository: trxTicketRepository, eventRepository: trxEventRepository, userRepository: trxUserRepository,
		ticketTransactionRepository: trxTicketTransactionRepository, usersAccessRepository: trxUsersAccessRepository}

}
