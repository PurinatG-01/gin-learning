package service

import (
	"errors"
	model "gin-learning/models"
	"gin-learning/repository"
	"time"
)

type EventService interface {
	All() (*[]model.Events, error)
	Create(model.FormEvent, int) (bool, error)
	List(page int, limit int) (model.Pagination[model.Events], error)
	Get(int) (model.Events, error)
	Delete(id int) (bool, error)
	Update(int, model.Events) (bool, error)
}

type eventService struct {
	eventRepository  repository.EventRepository
	userRepository   repository.UserRepository
	ticketRepository repository.TicketRepository
}

func NewEventService(eventRepository repository.EventRepository, userRepository repository.UserRepository, ticketRepository repository.TicketRepository) EventService {
	return &eventService{eventRepository: eventRepository, userRepository: userRepository, ticketRepository: ticketRepository}
}

func (s *eventService) All() (*[]model.Events, error) {
	events, err := s.eventRepository.All()
	return events, err
}

func (s *eventService) Create(form_event model.FormEvent, userId int) (bool, error) {
	isAdmin, admin_err := s.userRepository.IsAdmin(userId)
	if admin_err != nil {
		return false, admin_err
	}
	if !isAdmin {
		return false, errors.New("Not admin")
	}
	event := s.MapFormEventToEvents(form_event)
	_, err := s.eventRepository.Create(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) List(page int, limit int) (model.Pagination[model.Events], error) {
	events_pagination, err := s.eventRepository.List(page, limit)
	return events_pagination, err
}

func (s *eventService) Get(id int) (model.Events, error) {
	event, err := s.eventRepository.Get(id)
	return event, err
}

func (s *eventService) Delete(id int) (bool, error) {
	event := model.Events{Id: id}
	_, err := s.eventRepository.Delete(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) Update(id int, event model.Events) (bool, error) {
	event.Id = id
	now := time.Now()
	event.UpdatedAt = &now
	_, err := s.eventRepository.Update(&event)
	if err != nil {
		return true, err
	}
	return true, nil
}

func (s *eventService) MapFormEventToEvents(form_event model.FormEvent) model.Events {
	return model.Events{
		Title:        form_event.Title,
		Description:  form_event.Description,
		StartedAt:    form_event.StartedAt,
		EndedAt:      form_event.EndedAt,
		ReleasedAt:   form_event.ReleasedAt,
		TotalTickets: form_event.TotalTickets,
		// AvailableTickets: &form_event.AvailableTickets,
		TicketPrice: form_event.TicketPrice,
	}
}
