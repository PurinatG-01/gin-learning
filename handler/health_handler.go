package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (s *HealthHandler) ServerCheck(c *gin.Context) {
	c.HTML(http.StatusOK, "health.tmpl", gin.H{
		"msg": "🚀 Currently running!!",
	})
}
