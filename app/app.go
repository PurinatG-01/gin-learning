package app

import (
	"context"
	"gin-learning/handler"
	"gin-learning/repository"
	"gin-learning/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ApplicationContext struct {
	Auth    *handler.AuthHandler
	Event   *handler.EventHandler
	Ticket  *handler.TicketHandler
	User    *handler.UserHandler
	Health  *handler.HealthHandler
	Utility *handler.UtilityHandler
	DB      *gorm.DB
}

func NewApp(ctx context.Context) (*ApplicationContext, error) {

	// #0 [Pre] Setup database
	db, err := repository.ConnectDatabase()
	if err != nil {
		panic(err)
	}

	// #1 Init Repositories
	userRepository := repository.NewUserRepository(db)
	ticketRepository := repository.NewTicketRepository(db)
	eventRepository := repository.NewEventRepository(db)
	ticketTransactionRepository := repository.NewTicketTransactionRepository(db)
	usersAccessRepository := repository.NewUsersAccessRepository(db)

	// #2 Init Services
	// #2.1 Init authen/jwt/user services
	userService := service.NewUserService(userRepository, ticketRepository)
	jwtService := service.NewJWTService()
	loginService := service.NewLoginService(userRepository)
	// #2.2 Init ticket service
	ticketService := service.NewTicketService(ticketRepository, eventRepository, userRepository, ticketTransactionRepository, usersAccessRepository)
	// #2.3 Init event service
	eventService := service.NewEventService(eventRepository, userRepository, ticketRepository)

	// #3 Init handler/controller
	// #3.1 Init auth handler
	authHandler := handler.NewAuthHandler(loginService, jwtService, userService)
	// #3.2 Init ticket handler
	ticketHandler := handler.NewTicketHandler(ticketService)
	// #3.3 Init event handler
	eventHandler := handler.NewEventHandler(eventService)
	// #3.4 Init user handler
	userHandler := handler.NewUserHandler(userService)
	// #3.4 Init utility handler
	utilityHandler := handler.NewUtilityHandler()
	// #3.5 Init health handler
	healthHandler := handler.NewHealthHandler()

	return &ApplicationContext{
		Auth:    authHandler,
		Event:   eventHandler,
		Ticket:  ticketHandler,
		User:    userHandler,
		Health:  healthHandler,
		Utility: utilityHandler,
		DB:      db,
	}, nil
}

func InitApp(ctx context.Context, engine *gin.Engine) {
	app, _ := NewApp(ctx)
	InitRoutes(ctx, engine, app)
}
