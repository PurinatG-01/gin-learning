package service

import (
	model "gin-learning/models"
	"gin-learning/repository"
)

type EventService interface {
	All() (*[]model.Event, error)
}

// TODO: update to have repository
type eventService struct {
	repository repository.EventRepository
}

func NewEventService(repository repository.EventRepository) EventService {
	return &eventService{repository: repository}
}

func (s *eventService) All() (*[]model.Event, error) {
	events, _ := s.repository.All()
	return events, nil
}
