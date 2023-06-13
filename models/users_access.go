package model

import "time"

type UsersAccess struct {
	Id        int        `gorm:"id" json:"id"`
	TicketId  string     `gorm:"ticket_id" json:"ticketId"`
	UserId    int        `gorm:"user_id" json:"userId"`
	EventId   int        `gorm:"event_id" json:"eventId"`
	CreatedAt *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updatedAt"`
}
