package service

import (
	"context"
	model "gin-learning/models"
)

type EventService interface {
	All(ctx context.Context) ([]model.Event, error)
}

// TODO: update to have repository
type eventService struct {
}

func NewEventService() EventService {
	return &eventService{}
}

func (s *eventService) All(ctx context.Context) ([]model.Event, error) {
	return []model.Event{}, nil
}
