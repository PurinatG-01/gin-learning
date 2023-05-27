package app

import (
	"context"
	"gin-learning/handler"
	"gin-learning/repository"
	"gin-learning/service"

	"github.com/gin-gonic/gin"
)

type ApplicationContext struct {
	Auth    *handler.AuthHandler
	Event   *handler.EventHandler
	Ticket  *handler.TicketHandler
	Health  *handler.HealthHandler
	Utility *handler.UtilityHandler
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {

	// Set up database
	db, err := repository.ConnectDatabase()

	if err != nil {
		panic(err)
	}

	// Init authen app
	email := "purinat.san@gmail.com"
	password := "1234"
	jwtService := service.NewJWTService()
	loginService := service.NewLoginService(email, password)
	authHandler := handler.NewAuthHandler(loginService, jwtService)

	// Init Event app
	eventRepository := repository.NewEventRepository(db)
	eventService := service.NewEventService(eventRepository)
	eventHandler := handler.NewEventHandler(eventService)

	// Init Ticket app
	ticketRepository := repository.NewTicketRepository(db)
	ticketService := service.NewTicketService(ticketRepository)
	ticketHandler := handler.NewTicketHandler(ticketService)

	// Init Health app
	healthHandler := handler.NewHealthHandler()

	// Utility app
	utilityHandler := handler.NewUtilityHandler()

	return &ApplicationContext{
		Auth:    authHandler,
		Event:   eventHandler,
		Ticket:  ticketHandler,
		Health:  healthHandler,
		Utility: utilityHandler,
	}, nil
}

func InitApp(ctx context.Context, engine *gin.Engine) {
	app, _ := NewApp(ctx)
	InitRoutes(ctx, engine, app)
}
