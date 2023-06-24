package repository

import (
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

// WithTrx enables repository with transaction
func (s ticketTransactionAdapter) WithTrx(trxHandle *gorm.DB) TicketTransactionRepository {
	if trxHandle == nil {
		log.Print("[TicketTransaction] Transaction Database not found")
		return &ticketTransactionAdapter{DB: trxHandle}
	}
	s.DB = trxHandle
	return &ticketTransactionAdapter{DB: trxHandle}
}
