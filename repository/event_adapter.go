package repository

import (
	"errors"
	"fmt"
	model "gin-learning/models"
	"log"

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

func (s *eventAdapter) GetByKey(key string, value string) (model.Events, error) {
	var eventStruct model.Events
	result := s.DB.Where(fmt.Sprintf("%s = ?", key), value).First(&eventStruct)
	if result.RowsAffected != 1 {
		return eventStruct, errors.New(fmt.Sprintf("%s : %s found more than 1 (rows affeceted more than 1)", key, value))
	} else if result.Error != nil {
		return eventStruct, result.Error
	}
	return eventStruct, nil
}

// WithTrx enables repository with transaction
func (s eventAdapter) WithTrx(trxHandle *gorm.DB) EventRepository {
	if trxHandle == nil {
		log.Print("[Event] Transaction Database not found")
		return &eventAdapter{DB: trxHandle}
	}
	s.DB = trxHandle
	return &eventAdapter{DB: trxHandle}
}
