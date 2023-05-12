package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiResponse struct {
	Status int            `json:"status"`
	Data   map[string]any `json:"data"`
}

func ResponseMapper(status int, data map[string]any) ApiResponse {
	return ApiResponse{Status: status, Data: data}
}

func InitRoutes(engine *gin.Engine) {

	// GET routes
	engine.GET("/", landing)
	engine.GET("/ping", ping)
	engine.GET("/test", test)

	// Feed Post group
	post := engine.Group("/post")
	{
		post.GET("/list", test)
		post.POST("/", test)
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
	body := ResponseMapper(http.StatusOK, data)
	c.JSON(http.StatusOK, &body)
}
