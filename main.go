package main

import (
	"gin-learning/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	logger := log.InitLog(log.Logger{Name: "LOG #1"})
	logger.Log("Sever opened ja")
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")
	initRoutes(engine)
	engine.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func initRoutes(engine *gin.Engine) {

	// GET routes
	engine.GET("/", landing)
	engine.GET("/ping", ping)
	engine.GET("/test", test)

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
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "🚀 success",
	})
}
