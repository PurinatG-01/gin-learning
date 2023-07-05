package service

import "net/http"

type PaymentService interface {
	PurchaseByPromptpay() (interface{}, error)
}

func NewPaymentService(client *http.Client) PaymentService {
	return &paymentService{httpClient: client}
}

type paymentService struct {
	httpClient *http.Client
}

func (s *paymentService) PurchaseByPromptpay() (interface{}, error) {

	return nil, nil
}
