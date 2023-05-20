package repository

import (
	"errors"
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

func (s *eventAdapter) Get(id int) (model.Event, error) {
	var event model.Event
	result := s.DB.First(&event, id)
	if result.RowsAffected == 0 {
		return event, errors.New("[GET] event id not found")
	}
	return event, result.Error
}

func (s *eventAdapter) Update(event *model.Event) (bool, error) {
	result := s.DB.Model(event).Updates(event)
	return true, result.Error
}

func (s *eventAdapter) Delete(event *model.Event) (bool, error) {
	result := s.DB.Delete(event)
	return true, result.Error
}
