package repository

import (
	model "gin-learning/models"

	"gorm.io/gorm"
)

func NewTicketTransactionRepository(db *gorm.DB) TicketTransactionRepository {
	return &ticketTransactionAdapter{DB: db}
}

type ticketTransactionAdapter struct {
	DB *gorm.DB
}

func (s *ticketTransactionAdapter) Create(ticketTransaction *model.TicketsTransaction) (bool, error) {
	return true, nil
}
