package app

import (
	"gin-learning/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func InitRoutes(engine *gin.Engine, app *ApplicationContext) {

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
	}

	ticket := engine.Group("/ticket")
	{
		ticket.POST("/purchase", middleware.UserAuthorizeJWT(), middleware.DBTransactionMiddleware(app.DB), app.Ticket.Purchase)
	}

	payment := engine.Group("/payment")
	{
		payment.GET("/channel", app.Purchase.AllPaymentMethod)
		payment.POST("/callback", middleware.OmiseMiddleware(), middleware.DBTransactionMiddleware(app.DB), app.Purchase.OmiseHook)
		purchase := payment.Group("/purchase", middleware.UserAuthorizeJWT(), middleware.DBTransactionMiddleware(app.DB))
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
			userAuthen.GET("/transactions", app.User.Transactions)
			userAuthen.PUT("/update", app.User.Update)
		}
	}

	utility := engine.Group("/utility")
	{
		utility.GET("/random", app.Utility.Random)
	}

}
