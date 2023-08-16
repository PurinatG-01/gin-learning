package main

import (
	"fmt"
	"gin-learning/app"
	_ "gin-learning/docs"
	"gin-learning/log"
	"gin-learning/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title Gin-learning-event API
// @version 1.0
// @description.markdown

// @contact.name API Support
// @contact.url https://github.com/PurinatG-01
// @contact.email purinat.san@gmail.com

// @schemes https http

// @securityDefinitions.apikey JWT
// @in header
// @name Authorization
func main() {
	godotenv.Load(".env")
	logger := log.InitLog(log.Logger{Name: "LOG #1"})
	logger.Log("Sever opened ja")
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	// Set up cors middleware
	middleware.InitMiddlewares(engine)
	// Set up server routes
	app.InitApp(engine)

	engine.Run(fmt.Sprintf(":8080")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
