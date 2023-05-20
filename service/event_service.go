package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
	"time"
)

type EventService interface {
	All() (*[]model.Event, error)
	Create(model.Event) (bool, error)
	Get(int) (model.Event, error)
	Delete(id int) (bool, error)
	Update(int, model.Event) (bool, error)
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

func (s *eventService) Get(id int) (model.Event, error) {
	event, err := s.repository.Get(id)
	return event, err
}

func (s *eventService) Delete(id int) (bool, error) {
	event := model.Event{Id: id}
	_, err := s.repository.Delete(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) Update(id int, event model.Event) (bool, error) {
	event.Id = id
	now := time.Now()
	event.UpdatedAt = &now
	_, err := s.repository.Update(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}
