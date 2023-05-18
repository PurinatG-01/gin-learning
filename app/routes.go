package app

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitApp(ctx context.Context, engine *gin.Engine) {
	app, _ := NewApp(ctx)
	InitRoutes(ctx, engine, app)
}

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
		event.POST("/", test)
		event.DELETE("/", test)
		event.PUT("/", test)

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
