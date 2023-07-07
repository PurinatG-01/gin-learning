package service

import (
	"gin-learning/config"
	"log"
	"net/http"

	"github.com/omise/omise-go"
	"github.com/omise/omise-go/operations"
)

type PaymentService interface {
	PurchaseByPromptpay() (interface{}, error)
	GetPaymentConfig() (*[]config.PaymentMethod, error)
	CreatePromptpayCharge(amount int) (*omise.Charge, error)
}

func NewPaymentService(config *config.PaymentConfig) PaymentService {
	service := &paymentService{config: config}
	service.omiseClient = service.initOmiseClient(config.OmiseConfig.PublicKey, config.OmiseConfig.SecretKey)
	return service
}

type paymentService struct {
	httpClient  *http.Client
	omiseClient *omise.Client
	config      *config.PaymentConfig
}

func (s *paymentService) initOmiseClient(pk string, sk string) *omise.Client {
	client, e := omise.NewClient(pk, sk)
	if e != nil {
		log.Fatal(e)
	}
	return client
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

func (s *paymentService) CreatePromptpayCharge(amount int) (*omise.Charge, error) {
	map_amount := int64(amount * 100)
	source, createSource := &omise.Source{}, &operations.CreateSource{
		Amount:   map_amount,
		Currency: "thb",
		Type:     "promptpay",
	}
	if e := s.omiseClient.Do(source, createSource); e != nil {
		panic(e)
	}
	charge, createCharge := &omise.Charge{}, &operations.CreateCharge{
		Amount:   map_amount,
		Currency: "thb",
		Source:   source.ID,
	}
	if e := s.omiseClient.Do(charge, createCharge); e != nil {
		panic(e)
	}
	return charge, nil
}
