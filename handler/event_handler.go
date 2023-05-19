package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EventHandler struct {
	service   service.EventService
	responder utils.Responder
}

func NewEventHandler(service service.EventService) *EventHandler {

	return &EventHandler{service: service, responder: utils.Responder{}}
}

func (s *EventHandler) All(c *gin.Context) {
	events, err := s.service.All()
	if err != nil {
		s.responder.ResponseError(c, err.Error())
	}
	data := map[string]interface{}{"list": events}
	s.responder.ResponseSuccess(c, &data)
}

func (s *EventHandler) Create(c *gin.Context) {
	var event model.Event
	if err := c.BindJSON(&event); err != nil {
		s.responder.ResponseError(c, err.Error())
	}
	_, err := s.service.Create(event)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
	}
	c.JSON(http.StatusCreated, nil)
	// s.responder.ResponseSuccess(nil)
}

func (s *EventHandler) Get(c *gin.Context) {
	events, _ := s.service.Get()
	data := map[string]interface{}{"list": events}
	s.responder.ResponseSuccess(c, &data)
}

func (s *EventHandler) Delete(c *gin.Context) {
	events, _ := s.service.Delete()
	data := map[string]interface{}{"list": events}
	s.responder.ResponseSuccess(c, &data)
}

func (s *EventHandler) Update(c *gin.Context) {
	events, _ := s.service.Update()
	data := map[string]interface{}{"list": events}
	s.responder.ResponseSuccess(c, &data)
}
