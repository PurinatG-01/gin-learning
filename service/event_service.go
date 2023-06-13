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
	eventRepository  repository.EventRepository
	userRepository   repository.UserRepository
	ticketRepository repository.TicketRepository
}

func NewEventService(eventRepository repository.EventRepository, userRepository repository.UserRepository, ticketRepository repository.TicketRepository) EventService {
	return &eventService{eventRepository: eventRepository, userRepository: userRepository, ticketRepository: ticketRepository}
}

func (s *eventService) All() (*[]model.Event, error) {
	events, err := s.eventRepository.All()
	return events, err
}

// TODO: update create event to
func (s *eventService) Create(event model.Event) (bool, error) {
	_, err := s.eventRepository.Create(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) Get(id int) (model.Event, error) {
	event, err := s.eventRepository.Get(id)
	return event, err
}

func (s *eventService) Delete(id int) (bool, error) {
	event := model.Event{Id: id}
	_, err := s.eventRepository.Delete(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) Update(id int, event model.Event) (bool, error) {
	event.Id = id
	now := time.Now()
	event.UpdatedAt = &now
	_, err := s.eventRepository.Update(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}
