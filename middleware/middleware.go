package middleware

import (
	"fmt"
	"gin-learning/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/omise/omise-go"
	"gorm.io/gorm"
)

func InitMiddlewares(engine *gin.Engine) {
	engine.Use(CORSMiddlewares())
}

func CORSMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		allow_host := os.Getenv("APP_HOST")
		log.Print("[HOST]", allow_host)
		c.Writer.Header().Set("Access-Control-Allow-Origin", allow_host)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		// Handle pre-flight (OPTIONS) requests
		if c.Request.Method == "OPTIONS" {
			// Respond with 200 status code (Success) for pre-flight requests
			c.Writer.WriteHeader(http.StatusOK)
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

// StatusInList -> checks if the given status is in the list
func StatusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

// DBTransactionMiddleware : to setup the database transaction middleware
func DBTransactionMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		txHandle := db.Begin()
		log.Print("\n[Transaction] beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				txHandle.Rollback()
			}
		}()

		c.Set("db_trx", txHandle)
		c.Next()

		if StatusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			log.Print("\n [Transaction] committing transactions")
			if err := txHandle.Commit().Error; err != nil {
				log.Print("\n [Transaction] trx commit error: ", err)
			}
		} else {
			log.Print("\n[Transaction] rolling back transaction due to status code: ", c.Writer.Status())
			txHandle.Rollback()
		}
	}
}

func OmiseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		omise.WebhookHTTPHandler(&OmiseEventHandler{context: c}).ServeHTTP(c.Writer, c.Request)
	}
}

type OmiseEventHandler struct {
	context *gin.Context
}

func (s *OmiseEventHandler) HandleEvent(w http.ResponseWriter, r *http.Request, event *omise.Event) {
	s.context.Set("omise-event", event)
}
