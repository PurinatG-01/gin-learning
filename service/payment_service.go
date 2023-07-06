package service

import (
	"gin-learning/config"
	"net/http"
)

type PaymentService interface {
	PurchaseByPromptpay() (interface{}, error)
	GetPaymentConfig() (*[]config.PaymentMethod, error)
}

func NewPaymentService(config *config.PaymentConfig) PaymentService {
	return &paymentService{config: config}
}

type paymentService struct {
	httpClient *http.Client
	config     *config.PaymentConfig
}

func (s *paymentService) PurchaseByPromptpay() (interface{}, error) {
	// TODO: Create charge by Promptpay flow
	// #1 Generate Omise charge
	// #2 Derived Charge response to type PromptpayData
	// #3 Create row in tickets_transaction with status pending
	return nil, nil
}

func (s *paymentService) GetPaymentConfig() (*[]config.PaymentMethod, error) {
	return &s.config.PaymentMethodList, nil
}
