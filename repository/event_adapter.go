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

func (s *eventAdapter) All() (*[]model.Events, error) {
	var events *[]model.Events
	result := s.DB.Find(&events)
	return events, result.Error
}

func (s *eventAdapter) Create(event *model.Events) (bool, error) {
	result := s.DB.Create(event)
	return true, result.Error
}

func (s *eventAdapter) Get(id int) (model.Events, error) {
	var event model.Events
	result := s.DB.First(&event, id)
	if result.RowsAffected == 0 {
		return event, errors.New("[GET] event id not found")
	}
	return event, result.Error
}

func (s *eventAdapter) Update(event *model.Events) (bool, error) {
	result := s.DB.Model(event).Updates(event)
	return true, result.Error
}

func (s *eventAdapter) Delete(event *model.Events) (bool, error) {
	result := s.DB.Delete(event)
	return true, result.Error
}
