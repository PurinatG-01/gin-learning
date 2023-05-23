package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRoutes(ctx context.Context, engine *gin.Engine, app *ApplicationContext) {

	// GET routes
	engine.GET("/", landing)
	engine.GET("/ping", ping)
	engine.GET("/test", test)

	// Health
	engine.GET("/health", app.Health.ServerCheck)

	// Feed Post group
	post := engine.Group("/post")
	{
		post.GET("/list", test)
		post.POST("/", test)
	}

	event := engine.Group("/event")
	{
		event.GET("/", app.Event.All)
		event.GET("/:id", app.Event.Get)
		event.POST("/", app.Event.Create)
		event.DELETE("/:id", app.Event.Delete)
		event.PUT("/:id", app.Event.Update)

	}

	ticket := engine.Group("/ticket")
	{
		ticket.GET("/", app.Ticket.All)
		ticket.POST("/", app.Ticket.Create)

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
