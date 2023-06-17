package model

import "time"

type Tickets struct {
	Id          int        `gorm:"id" json:"id"`
	Price       int        `gorm:"price" json:"price"`
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
