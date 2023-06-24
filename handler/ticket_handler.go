package handler

import (
	"encoding/json"
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (s *TicketHandler) Purchase(c *gin.Context) {
	// #1 Get userId, eventId
	str_userId := c.GetString("x-user-id")
	userId, conv_err := strconv.Atoi(str_userId)
	var form_ticket model.FormTicket
	if conv_err != nil {
		s.responder.ResponseError(c, conv_err.Error())
		return
	}
	bind_err := c.ShouldBind(&form_ticket)
	if bind_err != nil {
		s.responder.ResponseError(c, bind_err.Error())
		return
	}
	// #2 Ticket.service purchase ticket
	txHandle := c.MustGet("db_trx").(*gorm.DB)
	ticket, purchase_err, is_serv_err := s.service.WithTrx(txHandle).Purchase(form_ticket, userId)
	// #2.1 Error (Money not enough, Create error etc.) => (400, 500)
	if purchase_err != nil {
		if is_serv_err {
			s.responder.ResponseServerError(c, purchase_err.Error())
			return
		} else {
			s.responder.ResponseError(c, purchase_err.Error())
			return
		}
	}
	// #2.2 Success => 200
	s.responder.ResponseSuccess(c, &map[string]interface{}{"ticket": ticket})
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
	_, err := s.service.Delete(param_id)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseUpdateSuccess(c)
	return
}

func (s *TicketHandler) Update(c *gin.Context) {
	param_id := c.Param("id")
	var ticket model.Tickets
	if err := c.BindJSON(&ticket); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	_, err := s.service.Update(param_id, ticket)
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	s.responder.ResponseUpdateSuccess(c)
	return
}
