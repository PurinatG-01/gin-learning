package service

// type PurchaseService interface {
// 	Purchase(form_ticket model.FormTicket, userId int) (is_success bool, err error, is_serv_err bool)
// 	MapFormTicketToTickets(form_ticket model.FormTicket, event model.Events, userId int) model.Tickets
// }

// func NewPurchaseService(ticketRepository repository.TicketRepository, eventRepository repository.EventRepository, userRepository repository.UserRepository) TicketService {
// 	return &purchaseService{ticketRepository: ticketRepository, eventRepository: eventRepository, userRepository: userRepository}
// }

// type purchaseService struct {
// 	ticketRepository repository.TicketRepository
// 	eventRepository  repository.EventRepository
// 	userRepository   repository.UserRepository
// }

// func (s *PurchaseService) Purchase(form_ticket model.FormTicket, userId int) (bool, error, bool) {
// 	// #1 Check if eventId has available ticket left
// 	event, event_err := s.eventRepository.Get(form_ticket.EventId)
// 	if event_err != nil {
// 		return false, event_err, true
// 	}
// 	pretty.Print(event.AvailableTickets < form_ticket.Amount)
// 	// #1.1 Out of ticket => error
// 	if event.AvailableTickets < form_ticket.Amount {
// 		return false, errors.New("Out of ticket"), false
// 	}
// 	// #1.2 Success => continue
// 	// #2 Check if userId has enough money in wallet
// 	user, user_err := s.userRepository.Get(userId)
// 	if user_err != nil {
// 		return false, user_err, true
// 	}
// 	// #2.1 Not enough money => error
// 	if user.TotalMoney < (event.TicketPrice * form_ticket.Amount) {
// 		return false, errors.New("Not have enough money"), false
// 	}
// 	// #2.2 Success => continue
// 	// #3 Create Ticket from total amount
// 	ticket := s.MapFormTicketToTickets(form_ticket, event, userId)
// 	_, ticket_err := s.ticketRepository.Create(&ticket)
// 	if ticket_err != nil {
// 		return false, ticket_err, true
// 	}
// 	// #4 Create Ticket transaction from total amount
// 	// #5 Create Ticket user access from total amount
// 	// #6 Update Event available tickets by counting from tickets

// 	return true, nil, false
// }

// func (s *ticketService) MapFormTicketToTickets(form_ticket model.FormTicket, event model.Events, userId int) model.Tickets {
// 	return model.Tickets{
// 		EventId: event.TicketPrice,
// 		OwnerId: userId,
// 	}
// }
