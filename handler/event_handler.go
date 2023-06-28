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

func (s *EventHandler) List(c *gin.Context) {
	str_page := c.Query("page")
	str_limit := c.Query("limit")
	var page, limit int
	if str_page == "" {
		page = 1
	} else {
		conv_page, page_err := strconv.Atoi(str_page)
		page = conv_page
		if page_err != nil {
			s.responder.ResponseError(c, page_err.Error())
			return
		}
	}
	if str_limit == "" {
		limit = 10
	} else {
		conv_limit, limit_err := strconv.Atoi(str_limit)
		limit = conv_limit
		if limit_err != nil {
			s.responder.ResponseError(c, limit_err.Error())
			return
		}
	}
	events, list_err := s.service.List(page, limit)
	if list_err != nil {
		s.responder.ResponseError(c, list_err.Error())
		return
	}
	marshal_events, _ := json.Marshal(&events)
	var data map[string]interface{}
	_ = json.Unmarshal(marshal_events, &data)
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
	marshal_event, _ := json.Marshal(&event)
	var data map[string]interface{}
	_ = json.Unmarshal(marshal_event, &data)
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
