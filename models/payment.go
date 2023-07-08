package model

const TICKET_PURCHASE_LIMIT = 5

type FormTicketPayment struct {
	EventId int    `form:"eventId" binding:"required"`
	Amount  int    `form:"amount" binding:"required"`
	Channel string `form:"channel" binding:"required"`
}
