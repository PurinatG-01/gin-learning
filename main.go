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
	engine.Use(CORSMiddleware())
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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
