package repository

import (
	"fmt"
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
	fmt.Printf("> [ALL] %v : \n", events)
	fmt.Printf("> [result] : %v \n", result.RowsAffected)
	fmt.Printf("> [result : error] : %v \n", result.Error)
	return events, nil
}
