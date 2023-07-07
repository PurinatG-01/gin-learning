package app

import (
	"context"
	"gin-learning/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRoutes(ctx context.Context, engine *gin.Engine, app *ApplicationContext) {

	// Swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Authntication routes
	engine.POST("/login", app.Auth.Login)
	engine.POST("/signup", app.Auth.Signup)

	// Health
	engine.GET("/health", app.Health.ServerCheck)

	event := engine.Group("/event")
	{
		event.GET("/", app.Event.All)
		event.GET("/list", app.Event.List)
		event.GET("/:id", app.Event.Get)
		event.Use(middleware.UserAuthorizeJWT()).POST("/", app.Event.Create)
		// event.DELETE("/:id", app.Event.Delete)
		// event.PUT("/:id", app.Event.Update)
	}

	ticket := engine.Group("/ticket")
	{
		// ticket.GET("/", app.Ticket.All)
		ticket.POST("/purchase", middleware.UserAuthorizeJWT(), middleware.DBTransactionMiddleware(app.DB), app.Ticket.Purchase)
		// ticket.POST("/", app.Ticket.Create)
		// ticket.GET("/:id", app.Ticket.Get)
		// ticket.DELETE("/:id", app.Ticket.Delete)
		// ticket.PUT("/:id", app.Ticket.Update)
	}

	payment := engine.Group("/payment")
	{
		payment.GET("/channel", app.Purchase.AllPaymentMethod)
		payment.POST("/hooks", app.Purchase.Test)
		payment.POST("/test", app.Purchase.TestCharge)
		purchase := payment.Group("/purchase", middleware.UserAuthorizeJWT())
		{
			purchase.POST("/ticket", app.Purchase.PurchaseTicket)
		}
	}

	user := engine.Group("/user")
	{
		user.GET("/:id", app.User.GetPublic)
		userAuthen := user.Use(middleware.UserAuthorizeJWT())
		{
			userAuthen.GET("/tickets", app.User.Tickets)
		}
	}

	utility := engine.Group("/utility")
	{
		utility.GET("/random", app.Utility.Random)
	}

}
