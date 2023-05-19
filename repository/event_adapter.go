package repository

import (
	model "gin-learning/models"

	"gorm.io/gorm"
)

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventAdapter{DB: db}

}

type eventAdapter struct {
	DB *gorm.DB
}

func (s *eventAdapter) All() (*[]model.Event, error) {
	var events *[]model.Event
	result := s.DB.Find(&events)
	return events, result.Error
}

func (s *eventAdapter) Create(event *model.Event) (bool, error) {
	result := s.DB.Create(event)
	return true, result.Error
}
