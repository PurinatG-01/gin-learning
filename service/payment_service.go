package service

import (
	"errors"
	"fmt"
	"gin-learning/config"
	model "gin-learning/models"
	"gin-learning/repository"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type PaymentService interface {
	PurchaseTicket(form_payment model.FormTicketPayment, user_id int) (*omise.Charge, error)
	GetPaymentConfig() (*[]config.PaymentMethod, error)
	CreatePromptpayCharge(charge *omise.Charge, amount int) error
	CheckPurchaseTicketQualification(form_payment model.FormTicketPayment) error
	ResolvePaymentChargeComplete(charge *omise.Charge) error
}

func NewPaymentService(eventRepository repository.EventRepository, ticketRepository repository.TicketRepository, ticketTransactionRepository repository.TicketTransactionRepository, usersAccessRepository repository.UsersAccessRepository, config *config.PaymentConfig) PaymentService {
	service := &paymentService{config: config, eventRepository: eventRepository, ticketRepository: ticketRepository, ticketTransactionRepository: ticketTransactionRepository, usersAccessRepository: usersAccessRepository}
	service.omiseClient = service.initOmiseClient(config.OmiseConfig.PublicKey, config.OmiseConfig.SecretKey)
	return service
}

type paymentService struct {
	httpClient                  *http.Client
	omiseClient                 *omise.Client
	config                      *config.PaymentConfig
	eventRepository             repository.EventRepository
	ticketRepository            repository.TicketRepository
	ticketTransactionRepository repository.TicketTransactionRepository
	usersAccessRepository       repository.UsersAccessRepository
}

func (s *paymentService) initOmiseClient(pk string, sk string) *omise.Client {
	client, e := omise.NewClient(pk, sk)
	if e != nil {
		log.Fatal(e)
	}
	return client
}

func (s *paymentService) CheckPurchaseTicketQualification(form_payment model.FormTicketPayment) error {
	// #1 Check if eventId has available ticket left
	// #1.1.0 Check if amount is not exceed
	if model.TICKET_PURCHASE_LIMIT < form_payment.Amount {
		return errors.New(fmt.Sprintf("Cannot purchase more than %d tickets", model.TICKET_PURCHASE_LIMIT))
	}
	// #1.1.1 Check if event exists
	event, event_err := s.eventRepository.Get(form_payment.EventId)
	if event_err != nil {
		return event_err
	}
	// #1.1.2 Check event available tickets by counting from tickets
	ticket_count, ticket_count_err := s.ticketRepository.Count(&model.Tickets{EventId: event.Id})
	if ticket_count_err != nil {
		return ticket_count_err
	}
	available_tickets := int(int64(event.TotalTickets) - ticket_count)
	// #1.1.3 Out of ticket => error
	if available_tickets < form_payment.Amount {
		return errors.New("Out of ticket")
	}
	// #1.2 Success => continue
	return nil
}

func (s *paymentService) PurchaseTicket(form_payment model.FormTicketPayment, user_id int) (*omise.Charge, error) {
	// #0 Get event
	event, event_err := s.eventRepository.Get(form_payment.EventId)
	if event_err != nil {
		return nil, event_err
	}
	// #0.1 Get total concurrent transactions on specific charge id (transaction id) (eventID, status = pending, status = successful)
	total_concurrent_transactions, count_err := s.ticketTransactionRepository.CountMultiple([]model.TicketsTransaction{
		{EventId: event.Id, Status: model.OMISE_CHARGE_STATUS_PENDING},
		{EventId: event.Id, Status: model.OMISE_CHARGE_STATUS_SUCCESSFUL},
	})
	if count_err != nil {
		return nil, count_err
	}
	// #0.2 Check if concurrent transactions sun with request amount exceed total tickets
	if total_concurrent_transactions+int64(form_payment.Amount) > int64(event.TotalTickets) {
		return nil, errors.New("Out of ticket")
	}

	// #1 Prepare omise charge
	var purchase_err error
	omise_charge := &omise.Charge{}
	// #1.1 [Promptpay] Create source & charge
	if form_payment.Channel == "promptpay" {
		charge_err := s.CreatePromptpayCharge(omise_charge, event.TicketPrice*form_payment.Amount)
		if charge_err != nil {
			purchase_err = charge_err
		}
	}
	if purchase_err != nil {
		return nil, purchase_err
	}

	// #2 Create ticket transaction with status pending along with charge ID
	ticket_transaction := model.TicketsTransaction{TicketId: nil, TransactionId: omise_charge.ID, PurchaserId: user_id, EventId: form_payment.EventId, Status: model.OMISE_CHARGE_STATUS_PENDING}
	ticket_transaction_list := []model.TicketsTransaction{}
	for i := 0; i < form_payment.Amount; i++ {
		ticket_transaction.Id = uuid.New().String()
		ticket_transaction_list = append(ticket_transaction_list, ticket_transaction)
	}
	_, ticket_transaction_err := s.ticketTransactionRepository.CreateMultiple(&ticket_transaction_list, 20)
	if ticket_transaction_err != nil {
		return nil, ticket_transaction_err
	}

	return omise_charge, nil
}

func (s *paymentService) CreatePromptpayCharge(charge *omise.Charge, amount int) error {
	map_amount := int64(amount * 100)
	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   map_amount,
		Currency: "thb",
		Type:     "promptpay",
	}
	if err := s.omiseClient.Do(source, createSource); err != nil {
		return err
	}
	createCharge := &operations.CreateCharge{
		Amount:   map_amount,
		Currency: "thb",
		Source:   source.ID,
	}
	if err := s.omiseClient.Do(charge, createCharge); err != nil {
		return err
	}
	return nil
}

func (s *paymentService) GetPaymentConfig() (*[]config.PaymentMethod, error) {
	return &s.config.PaymentMethodList, nil
}

func (s *paymentService) ValidateCharge(charge *omise.Charge) (*model.TicketsTransaction, error) {
	transaction, transaction_err := s.ticketTransactionRepository.GetByKey("transaction_id", charge.ID)
	if transaction_err != nil {
		return nil, errors.New("Transaction not found")
	}
	return &transaction, nil
}

func (s *paymentService) ResolvePaymentChargeComplete(charge *omise.Charge) error {
	// #0 Validate charge id
	transaction, transaction_err := s.ValidateCharge(charge)
	if transaction_err != nil {
		return transaction_err
	}
	tr_count, tr_count_err := s.ticketTransactionRepository.Count(&model.TicketsTransaction{TransactionId: charge.ID})
	if tr_count_err != nil {
		return tr_count_err
	}
	// #1 Update ticket transaction status to complete
	_, u_t_err := s.ticketTransactionRepository.UpdateByKey("transaction_id", transaction.Id, "status", model.OMISE_CHARGE_STATUS_SUCCESSFUL)
	if u_t_err != nil {
		return u_t_err
	}
	// #2 Create ticket
	ticket := model.Tickets{
		EventId: transaction.EventId,
		OwnerId: transaction.PurchaserId,
	}
	ticket_list := []model.Tickets{}
	for i := 0; i < int(tr_count); i++ {
		ticket.Id = uuid.New().String()
		ticket_list = append(ticket_list, ticket)
	}
	ticket_list, cr_t_err := s.ticketRepository.CreateMultiple(&ticket_list, 20)
	if cr_t_err != nil {
		return cr_t_err
	}
	// #3 Create Users access
	users_access := model.UsersAccess{UserId: transaction.PurchaserId, EventId: transaction.EventId}
	users_access_list := []model.UsersAccess{}
	for i := 0; i < int(tr_count); i++ {
		users_access.TicketId = ticket_list[i].Id
		users_access_list = append(users_access_list, users_access)
	}
	_, users_access_err := s.usersAccessRepository.CreateMultiple(&users_access_list, 20)
	if users_access_err != nil {
		return users_access_err
	}
	return nil
}
