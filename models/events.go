package model

import "time"

type Event struct {
	Id          int        `gorm:"id" json:"id"`
	Title       string     `gorm:"title" json:"title"`
	Description string     `gorm:"description" json:"description"`
	StartedAt   *time.Time `gorm:"started_at" json:"startedAt"`
	EndedAt     *time.Time `gorm:"ended_at" json:"endedAt"`
	ReleasedAt  *time.Time `gorm:"released_at" json:"releasedAt"`
	CreatedAt   *time.Time `gorm:"created_at" json:"createdAt"`
}
