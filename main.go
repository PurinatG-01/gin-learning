package main

import (
	"context"
	"gin-learning/app"
	"gin-learning/log"
	"gin-learning/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	logger := log.InitLog(log.Logger{Name: "LOG #1"})
	logger.Log("Sever opened ja")
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	ctx := context.Background()
	// Set up cors middleware
	middleware.InitMiddlewares(engine)
	// Set up server routes
	app.InitApp(ctx, engine)

	engine.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
