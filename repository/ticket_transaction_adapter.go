package repository

import (
	"errors"
	"fmt"
	model "gin-learning/models"
	"log"

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

func (s *ticketTransactionAdapter) Count(transaction *model.TicketsTransaction) (int64, error) {
	var count int64
	result := s.DB.Model(&model.TicketsTransaction{}).Where(transaction).Count(&count)
	return count, result.Error
}

func (s *ticketTransactionAdapter) GetByKey(key string, value string) (model.TicketsTransaction, error) {
	var transaction model.TicketsTransaction
	result := s.DB.Where(fmt.Sprintf("%s = ?", key), value).First(&transaction)
	if result.RowsAffected != 1 {
		return transaction, errors.New(fmt.Sprintf("%s : %s found more than 1 (rows affeceted more than 1)", key, value))
	} else if result.Error != nil {
		return transaction, result.Error
	}
	return transaction, nil
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
