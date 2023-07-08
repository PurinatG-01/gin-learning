package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const (
	OMISE_CHARGE_STATUS_PENDING    = "pending"
	OMISE_CHARGE_STATUS_FAILED     = "failed"
	OMISE_CHARGE_STATUS_SUCCESSFUL = "successful"
	OMISE_CHARGE_STATUS_REVERSED   = "reversed"
	OMISE_CHARFE_STATUS_EXPIRED    = "expired"
)

type TicketsTransaction struct {
	Id            string     `gorm:"id" json:"id"`
	TicketId      *string    `gorm:"ticket_id" json:"ticketId"`
	PurchaserId   int        `gorm:"purchaser_id" json:"purchaserId"`
	EventId       int        `gorm:"event_id" json:"eventId"`
	CreatedAt     *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `gorm:"updated_at" json:"updatedAt"`
	TransactionId string     `gorm:"transaction_id" json:"transactionId"`
	Status        string     `gorm:"status" json:"status"`
}

func (s *TicketsTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New().String()
	return
}
