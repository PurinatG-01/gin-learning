package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Tickets struct {
	Id          string     `gorm:"id" json:"id"`
	EventId     int        `gorm:"event_id" json:"eventId"`
	CreatedAt   *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt   *time.Time `gorm:"updated_at" json:"updatedAt"`
	OwnerId     int        `gorm:"owner_id" json:"ownerId"`
	PurchasedAt *time.Time `gorm:"purchasedAt" json:"purchasedAt"`
}

type FormTicket struct {
	EventId int `gorm:"event_id" form:"eventId" binding:"required"`
	Amount  int `form:"amount" binding:"required"`
}

func (s *Tickets) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	s.Id = uuid.New().String()
	s.PurchasedAt = &now
	return
}
