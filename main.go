package main

import (
	"gin-learning/db"
	"gin-learning/log"
	"gin-learning/middleware"
	"gin-learning/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	logger := log.InitLog(log.Logger{Name: "LOG #1"})
	logger.Log("Sever opened ja")
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	// Set up cors middleware
	middleware.InitMiddlewares(engine)
	// Set up server routes
	routes.InitRoutes(engine)
	// Set up database
	db.ConnectDatabase()

	engine.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
