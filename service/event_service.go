package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
)

type EventService interface {
	All() (*[]model.Event, error)
	Create(model.Event) (bool, error)
	Get() (model.Event, error)
	Delete() (bool, error)
	Update() (bool, error)
}

type eventService struct {
	repository repository.EventRepository
}

func NewEventService(repository repository.EventRepository) EventService {
	return &eventService{repository: repository}
}

func (s *eventService) All() (*[]model.Event, error) {
	events, err := s.repository.All()
	return events, err
}

func (s *eventService) Create(event model.Event) (bool, error) {
	_, err := s.repository.Create(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) Get() (model.Event, error) {
	// events, _ := s.repository.All()
	return model.Event{}, nil
}

func (s *eventService) Delete() (bool, error) {
	// events, _ := s.repository.All()
	return true, nil
}

func (s *eventService) Update() (bool, error) {
	// events, _ := s.repository.All()
	return true, nil
}
