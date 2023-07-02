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

// HealthHandler godoc
// @tags Health
// @summary Health Check
// @description Health checking for the service
// @id HealthServerCheckHandler
// @produce html
// @router /health [get]
func (s *HealthHandler) ServerCheck(c *gin.Context) {
	c.HTML(http.StatusOK, "health.tmpl", gin.H{
		"msg": "🚀 Currently running!!",
	})
}
