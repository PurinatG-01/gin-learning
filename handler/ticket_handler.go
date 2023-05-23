package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"

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
	var ticket model.Ticket
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

}

func (s *TicketHandler) Delete(c *gin.Context) {

}

func (s *TicketHandler) Update(c *gin.Context) {

}
