package repository

import (
	"errors"
	"fmt"
	model "gin-learning/models"
	"log"

	"github.com/omise/omise-go"
	"gorm.io/gorm"
)

func NewTicketTransactionRepository(db *gorm.DB) TicketTransactionRepository {
	return &ticketTransactionAdapter{DB: db}
}

type ticketTransactionAdapter struct {
	DB *gorm.DB
}

func (s *ticketTransactionAdapter) Create(ticketTransaction *model.TicketsTransaction) (model.TicketsTransaction, error) {
	result := s.DB.Create(ticketTransaction)
	return *ticketTransaction, result.Error
}

func (s *ticketTransactionAdapter) CreateMultiple(transactionList *[]model.TicketsTransaction, batchSize int) ([]model.TicketsTransaction, error) {
	result := s.DB.CreateInBatches(transactionList, batchSize)
	return *transactionList, result.Error
}

func (s *ticketTransactionAdapter) Get(id int) (model.TicketsTransaction, error) {
	var transaction model.TicketsTransaction
	result := s.DB.First(&transaction, id)
	if result.RowsAffected == 0 {
		return transaction, errors.New("[GET] ticket transaction id not found")
	}
	return transaction, result.Error
}

func (s *ticketTransactionAdapter) ListByUserId(userId int, page int, limit int, status omise.ChargeStatus, order model.OrderBy) (model.Pagination[model.TicketsTransaction], error) {
	var ticket_transaction_list *[]model.TicketsTransaction
	pagination := model.Pagination[model.TicketsTransaction]{Page: page, Limit: limit, Sort: fmt.Sprintf("created_at %s", order)}
	result := s.DB.Preload("Event")
	if order == "" {
		result = result.Order("created_at desc")
	} else {
		result = result.Order(fmt.Sprintf("created_at %s", order))
	}
	if status == "" {
		result = result.Find(&ticket_transaction_list, "purchaser_id = ?", userId)
	} else {
		result = result.Find(&ticket_transaction_list, "purchaser_id = ? AND status = ?", userId, status)
	}
	result = result.Scopes(Paginate(ticket_transaction_list, &pagination, result))
	pagination.List = *ticket_transaction_list
	return pagination, result.Error
}

func (s *ticketTransactionAdapter) Count(transaction *model.TicketsTransaction) (int64, error) {
	var count int64
	result := s.DB.Model(&model.TicketsTransaction{}).Where(transaction).Count(&count)
	return count, result.Error
}

func (s *ticketTransactionAdapter) CountMultiple(list []model.TicketsTransaction) (int64, error) {
	var count int64
	result := s.DB.Model(list).Count(&count)
	return count, result.Error
}

func (s *ticketTransactionAdapter) CountFromEventIdAndStatus(eventId int, statusList []string) (int64, error) {
	var count int64
	result := s.DB.Model(&model.TicketsTransaction{}).Where("event_id = ? AND status IN ?", eventId, statusList).Count(&count)
	return count, result.Error
}

func (s *ticketTransactionAdapter) GetByKey(key string, value string) (model.TicketsTransaction, error, int) {
	var transaction model.TicketsTransaction
	result := s.DB.Where(fmt.Sprintf("%s = ?", key), value).First(&transaction)
	if result.RowsAffected != 1 {
		return transaction, errors.New(fmt.Sprintf("%s : %s found more than 1 (rows affeceted more than 1)", key, value)), int(result.RowsAffected)
	} else if result.Error != nil {
		return transaction, result.Error, int(result.RowsAffected)
	}
	return transaction, nil, int(result.RowsAffected)
}

func (s *ticketTransactionAdapter) UpdateByKey(fkey string, fvalue any, skey string, svalue any) (bool, error) {
	result := s.DB.Model(&model.TicketsTransaction{}).Where(fmt.Sprintf("%s = ?", fkey), fvalue).Updates(map[string]interface{}{skey: svalue})
	return true, result.Error
}

// WithTrx enables repository with transaction
func (s ticketTransactionAdapter) WithTrx(trxHandle *gorm.DB) TicketTransactionRepository {
	if trxHandle == nil {
		log.Print("[TicketTransaction] Transaction Database not found")
		return &ticketTransactionAdapter{DB: trxHandle}
	}
	s.DB = trxHandle
	return &ticketTransactionAdapter{DB: trxHandle}
}
