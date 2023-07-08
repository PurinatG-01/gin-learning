package handler

import (
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	paymentService service.PaymentService
	ticketService  service.TicketService
	discordService service.DiscordService
	responder      utils.Responder
}

func NewPurchaseHandler(paymentService service.PaymentService, ticketService service.TicketService, discordService service.DiscordService) *PurchaseHandler {
	return &PurchaseHandler{
		paymentService: paymentService,
		ticketService:  ticketService,
		discordService: discordService,
		responder:      utils.Responder{},
	}
}

func (s *PurchaseHandler) PurchaseTicket(c *gin.Context) {
	str_userId := c.GetString("x-user-id")
	user_id, conv_err := strconv.Atoi(str_userId)
	if conv_err != nil {
		s.responder.ResponseError(c, conv_err.Error())
		return
	}
	var form_payment model.FormTicketPayment
	if err := c.ShouldBind(&form_payment); err != nil {
		s.responder.ResponseError(c, err.Error())
		return
	}
	if ql_err := s.paymentService.CheckPurchaseTicketQualification(form_payment); ql_err != nil {
		s.responder.ResponseError(c, ql_err.Error())
		return
	}

	charge, charge_err := s.paymentService.PurchaseTicket(form_payment, user_id)
	if charge_err != nil {
		s.responder.ResponseError(c, charge_err.Error())
		return
	}
	s.discordService.SendTransactionMessage(charge.ID, form_payment.Amount, user_id, form_payment.EventId, model.OMISE_CHARGE_STATUS_PENDING)
	s.responder.ResponseSuccess(c, &map[string]interface{}{"charge": charge})
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

func (s *PurchaseHandler) OmiseHook(c *gin.Context) {
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
