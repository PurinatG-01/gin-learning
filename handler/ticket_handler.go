package handler

import (
	"encoding/json"
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	service   service.TicketService
	responder utils.Responder
}

func NewTicketHandler(service service.TicketService) *TicketHandler {
	return &TicketHandler{service: service, responder: utils.Responder{}}
}

func (s *TicketHandler) All(c *gin.Context) {
	tickets, err := s.service.All()
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	data := map[string]interface{}{"list": tickets}
	s.responder.ResponseSuccess(c, &data)
	return
}

func (s *TicketHandler) Create(c *gin.Context) {
	var ticket model.Tickets
	if err := c.BindJSON(&ticket); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	_, err := s.service.Create(ticket)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseCreateSuccess(c)
	return
}

func (s *TicketHandler) Get(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	ticket, err := s.service.Get(id)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	b, _ := json.Marshal(&ticket)
	var data map[string]interface{}
	_ = json.Unmarshal(b, &data)
	s.responder.ResponseSuccess(c, &data)
	return

}

func (s *TicketHandler) Delete(c *gin.Context) {
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

func (s *TicketHandler) Update(c *gin.Context) {
	param_id := c.Param("id")
	id, param_err := strconv.Atoi(param_id)
	if param_err != nil {
		s.responder.ResponseError(c, param_err.Error())
		return
	}
	var ticket model.Tickets
	if err := c.BindJSON(&ticket); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	_, err := s.service.Update(id, ticket)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseUpdateSuccess(c)
	return
}
