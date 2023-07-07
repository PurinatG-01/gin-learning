package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"

	"github.com/gin-gonic/gin"
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

	s.responder.ResponseSuccess(c, &map[string]interface{}{"acknowledged": true})
	return
}

func (s *PurchaseHandler) AllPaymentMethod(c *gin.Context) {
	payment_methods, err := s.paymentService.GetPaymentConfig()
	if err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	data := map[string]interface{}{"list": payment_methods}
	s.responder.ResponseSuccess(c, &data)
	return
}

func (s *PurchaseHandler) TestCharge(c *gin.Context) {
	charge, charge_err := s.paymentService.CreatePromptpayCharge(525)
	data := map[string]interface{}{"charge_err": charge_err, "charge": charge}
	s.responder.ResponseSuccess(c, &data)
	return
}

func (s *PurchaseHandler) Test(c *gin.Context) {
	var data map[string]interface{}
	if err := c.BindJSON(&data); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	// pretty.Print("----------- \n")
	// pretty.Print(c.Request.Header)
	// pretty.Print("----------- \n")

	// pretty.Print(data)
	s.responder.ResponseSuccess(c, &data)
	return
}
