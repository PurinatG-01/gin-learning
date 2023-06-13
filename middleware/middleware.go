package middleware

import (
	"fmt"
	"gin-learning/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitMiddlewares(engine *gin.Engine) {
	engine.Use(CORSMiddlewares())
}

func CORSMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("APP_HOST"))
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

func UserAuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) <= 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, _, claims := service.NewJWTService().ValidateToken(tokenString)
		userId := claims["id"].(float64)
		username := claims["username"].(string)
		// Set context
		c.Set("x-user-id", fmt.Sprintf("%v", userId))
		c.Set("x-username", fmt.Sprintf("%v", username))
		if token.Valid {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
