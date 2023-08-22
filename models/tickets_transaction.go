package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/omise/omise-go"
	"gorm.io/gorm"
)

type OmiseChargeStatus = omise.ChargeStatus

const (
	OMISE_CURRENCY_RATE_TH                           = 100
	OMISE_CHARGE_SCOPE             omise.SearchScope = omise.ChargeScope
	OMISE_CHARGE_STATUS_ALL        OmiseChargeStatus = "all"
	OMISE_CHARGE_STATUS_PENDING    OmiseChargeStatus = omise.ChargePending
	OMISE_CHARGE_STATUS_FAILED     OmiseChargeStatus = omise.ChargeFailed
	OMISE_CHARGE_STATUS_SUCCESSFUL OmiseChargeStatus = omise.ChargeSuccessful
	OMISE_CHARGE_STATUS_REVERSED   OmiseChargeStatus = omise.ChargeReversed
	OMISE_CHARGE_STATUS_EXPIRED    string            = "expired"
)

type TicketsTransaction struct {
	Id            string     `gorm:"id" json:"id"`
	TicketId      *string    `gorm:"ticket_id" json:"ticketId"`
	PurchaserId   int        `gorm:"purchaser_id" json:"purchaserId"`
	EventId       int        `gorm:"event_id" json:"eventId"`
	CreatedAt     *time.Time `gorm:"created_at" json:"createdAt"`
	UpdatedAt     *time.Time `gorm:"updated_at" json:"updatedAt"`
	TransactionId string     `gorm:"transaction_id" json:"transactionId"`
	Status        string     `gorm:"status" json:"status"`
	Event         Events     `json:"event"`
}

func (s *TicketsTransaction) BeforeCreate(tx *gorm.DB) (err error) {
	s.Id = uuid.New().String()
	return
}

type FormTicketTransactionList struct {
	OrderBy OrderBy           `form:"orderBy"`
	Status  OmiseChargeStatus `form:"status"`
}
