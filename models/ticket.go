package model

import "time"

type Ticket struct {
	Id        int        `gorm:"id" json:"id"`
	Price     int        `gorm:"price" json:"price"`
	CreatedAt *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"updated_at" json:"updatedAt"`
}
