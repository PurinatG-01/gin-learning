package handler

import (
	"gin-learning/service"
	"gin-learning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service service.EventService
}

func NewEventHandler(service service.EventService) *EventHandler {
	return &EventHandler{service: service}
}

func (s *EventHandler) Test(c *gin.Context) {
	data := map[string]any{"msg": "🚀 KRUBB"}
	c.JSON(http.StatusOK, utils.ResponseMapper(http.StatusOK, &data))
}
