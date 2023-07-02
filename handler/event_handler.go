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
	paginator utils.Paginator
}

func NewEventHandler(service service.EventService) *EventHandler {

	return &EventHandler{service: service, responder: utils.Responder{}, paginator: utils.Paginator{}}
}

// AllEvent godoc
// @summary All Events
// @description All event
// @tags Events
// @id EventAllHandler
// @produce json
// @response 200 {object} utils.ApiResponse
// @Router /event [get]
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

// ListEvent godoc
// @summary List Events
// @description List event by pagination
// @tags Events
// @id EventListHandler
// @produce json
// @param page query int true "page of the list"
// @param limit query int true "limit of the list"
// @response 200 {object} utils.ApiResponse
// @Router /event/list [get]
func (s *EventHandler) List(c *gin.Context) {
	paginator_err := s.paginator.Bind(c)
	if paginator_err != nil {
		s.responder.ResponseError(c, paginator_err.Error())
		return
	}
	events, list_err := s.service.List(s.paginator.Page, s.paginator.Limit)
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

// CreateEvent godoc
// @summary Create Event
// @description Create event
// @security JWT
// @tags Events
// @id EventCreateHandler
// @accept mpfd
// @produce json
// @param body formData model.FormEvent true "Event data to be create"
// @response 201 {object} utils.ApiResponse
// @Router /event [post]
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

// GetEvent godoc
// @summary Get Event
// @description Get event by id
// @tags Events
// @id EventGetHandler
// @produce json
// @param id path int true "Event ID"
// @response 200 {object} utils.ApiResponse
// @Router /event/{id}  [get]
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
