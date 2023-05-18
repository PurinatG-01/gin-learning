package app

import (
	"context"
	"gin-learning/handler"
	"gin-learning/service"
)

type ApplicationContext struct {
	Event  *handler.EventHandler
	Health *handler.HealthHandler
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {

	eventService := service.NewEventService()
	eventHandler := handler.NewEventHandler(eventService)

	healthHandler := handler.NewHealthHandler()

	return &ApplicationContext{
		Event:  eventHandler,
		Health: healthHandler,
	}, nil
}
