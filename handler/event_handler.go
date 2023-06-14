package handler

import (
	"encoding/json"
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

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
		return
	}
	data := map[string]interface{}{"list": events}
	s.responder.ResponseSuccess(c, &data)
	return
}

func (s *EventHandler) Create(c *gin.Context) {
	userId, conv_err := strconv.Atoi(c.GetString("x-user-id"))
	if conv_err != nil {
		s.responder.ResponseError(c, conv_err.Error())
		return
	}

	// Binding form event
	var form_event model.FormEvent
	bind_err := c.ShouldBind(&form_event)
	if bind_err != nil {
		s.responder.ResponseError(c, bind_err.Error())
		return
	}

	// Create Event
	_, err := s.service.Create(form_event, userId)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseCreateSuccess(c)
	return
}

func (s *EventHandler) Get(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	event, err := s.service.Get(id)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	b, _ := json.Marshal(&event)
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	s.responder.ResponseSuccess(c, &data)
	return
}

func (s *EventHandler) Delete(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	_, err := s.service.Delete(id)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseUpdateSuccess(c)
	return
}

func (s *EventHandler) Update(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	var event model.Events
	if err := c.BindJSON(&event); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	_, err := s.service.Update(id, event)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseUpdateSuccess(c)
	return
}
