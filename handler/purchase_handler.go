package handler

import (
	"errors"
	model "gin-learning/models"
	"gin-learning/service"
	"gin-learning/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/omise/omise-go"
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
	event, exists := c.Get("omise-event")
	if !exists {
		s.responder.ResponseError(c, errors.New("event not found").Error())
		return
	}
	event_data := event.(*omise.Event)
	if event_data.Key == "charge.complete" {
		charge := event_data.Data.(*omise.Charge)
		if charge.Status == model.OMISE_CHARGE_STATUS_SUCCESSFUL {
			resolve_err := s.paymentService.ResolvePaymentChargeComplete(charge)
			if resolve_err != nil {
				s.responder.ResponseError(c, resolve_err.Error())
				return
			}
			s.discordService.SendTransactionMessage(charge.ID, int(charge.Amount), string(model.OMISE_CHARGE_STATUS_SUCCESSFUL))
		} else if charge.Status == model.OMISE_CHARGE_STATUS_FAILED {
			resolve_err := s.paymentService.ResolvePaymentChargeFailed(charge)
			if resolve_err != nil {
				s.responder.ResponseError(c, resolve_err.Error())
				return
			}
			s.discordService.SendTransactionMessage(charge.ID, int(charge.Amount), string(model.OMISE_CHARGE_STATUS_FAILED))
		}
	} else if event_data.Key == "charge.create" {
		charge := event_data.Data.(*omise.Charge)
		s.discordService.SendTransactionMessage(charge.ID, int(charge.Amount), string(model.OMISE_CHARGE_STATUS_PENDING))
	}
	s.responder.ResponseSuccess(c, &map[string]interface{}{"data": event_data})
	return
}
