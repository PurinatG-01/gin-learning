package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TicketsTransaction struct {
	Id          string     `gorm:"id" json:"id"`
	TicketId    string     `gorm:"ticket_id" json:"ticketId"`
	PurchaserId int        `gorm:"purchaser_id" json:"purchaserId"`
	EventId     int        `gorm:"event_id" json:"eventId"`
	CreatedAt   *time.Time `gorm:"created_at" json:"createdAt"`
}

func (s *TicketsTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New().String()
	return
}
