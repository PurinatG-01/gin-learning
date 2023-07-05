package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"

	"github.com/gin-gonic/gin"
	"github.com/kr/pretty"
)

type PurchaseHandler struct {
	paymentService service.PaymentService
	ticketService  service.TicketService
	responder      utils.Responder
}

func NewPurchaseHandler(paymentService service.PaymentService, ticketService service.TicketService) *PurchaseHandler {
	return &PurchaseHandler{
		paymentService: paymentService,
		ticketService:  ticketService,
		responder:      utils.Responder{},
	}
}

func (s *PurchaseHandler) PurchaseTicket(c *gin.Context) {
	var form_payment model.FormTicketPayment
	if err := c.ShouldBind(&form_payment); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}

	pretty.Print(form_payment)

	s.responder.ResponseSuccess(c, &map[string]interface{}{"acknowledged": true})
	return
}
