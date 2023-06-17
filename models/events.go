package model

import "time"

type Events struct {
	Id               int        `gorm:"id" json:"id"`
	Title            string     `gorm:"title" json:"title"`
	Description      string     `gorm:"description" json:"description"`
	StartedAt        *time.Time `gorm:"started_at" json:"startedAt"`
	EndedAt          *time.Time `gorm:"ended_at" json:"endedAt"`
	ReleasedAt       *time.Time `gorm:"released_at" json:"releasedAt"`
	CreatedAt        *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt        *time.Time `gorm:"updated_at" json:"updatedAt"`
	TotalTickets     int        `gorm:"total_tickets" json:"totalTickets"`
	AvailableTickets int        `gorm:"available_tickets" json:"availableTickets"`
	TicketPrice      int        `gorm:"ticket_price" json:"ticketPrice"`
}

type FormEvent struct {
	Title            string     `form:"title"`
	Description      string     `form:"description"`
	StartedAt        *time.Time `form:"startedAt"`
	EndedAt          *time.Time `form:"endedAt"`
	ReleasedAt       *time.Time `form:"releasedAt"`
	TotalTickets     int        `form:"totalTickets"`
	AvailableTickets int        `form:"availableTickets"`
	TicketPrice      int        `form:"ticketPrice"`
}
