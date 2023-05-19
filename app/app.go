package app

import (
	"context"
	"gin-learning/handler"
	"gin-learning/repository"
	"gin-learning/service"

	"github.com/gin-gonic/gin"
)

type ApplicationContext struct {
	Event  *handler.EventHandler
	Health *handler.HealthHandler
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {

	// Set up database
	db, err := repository.ConnectDatabase()

	if err != nil {
		panic("[APP] failed to connect database")
	}

	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)

	healthHandler := handler.NewHealthHandler()

	return &ApplicationContext{
		Event:  eventHandler,
		Health: healthHandler,
	}, nil
}

func InitApp(ctx context.Context, engine *gin.Engine) {
	app, _ := NewApp(ctx)
	InitRoutes(ctx, engine, app)
}
