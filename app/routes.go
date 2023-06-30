package app

import (
	"context"
	"gin-learning/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(ctx context.Context, engine *gin.Engine, app *ApplicationContext) {

	// GET routes
	engine.GET("/", landing)
	engine.GET("/ping", ping)
	engine.GET("/test", test)

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
		ticket.GET("/", app.Ticket.All)
		ticket.Use().Use().POST("/purchase", middleware.UserAuthorizeJWT(), middleware.DBTransactionMiddleware(app.DB), app.Ticket.Purchase)
		ticket.POST("/", app.Ticket.Create)
		ticket.GET("/:id", app.Ticket.Get)
		ticket.DELETE("/:id", app.Ticket.Delete)
		ticket.PUT("/:id", app.Ticket.Update)

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
		utility.GET("/shuffle", app.Utility.Shuffle)
	}

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func landing(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "Gin-learning",
	})
}

func test(c *gin.Context) {
	data := map[string]any{"msg": "🚀 Yeahhh", "list": []int{1, 2, 3, 4, 5}}
	c.JSON(http.StatusOK, &data)
}
