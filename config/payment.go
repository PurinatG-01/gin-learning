package config

import (
	"errors"
	"fmt"
	"os"
)

type PaymentConfig struct {
	PaymentMethodList []PaymentMethod
	OmiseConfig       OmiseConfig
}

type PaymentMethod struct {
	Name      string
	OmiseType string
}

type OmiseConfig struct {
	SecretKey string
	PublicKey string
}

func NewPaymentConfig() (*PaymentConfig, error) {
	payment := PaymentConfig{}
	pk := fmt.Sprintf("%s", os.Getenv("OMISE_PUBLIC_KEY"))
	sk := fmt.Sprintf("%s", os.Getenv("OMISE_SECRET_KEY"))
	if pk == "" || sk == "" {
		return nil, errors.New("omise key not found")
	}
	payment.OmiseConfig = OmiseConfig{SecretKey: sk, PublicKey: pk}
	payment.PaymentMethodList = []PaymentMethod{
		{Name: "Promptpay", OmiseType: "promptpay"},
	}
	return &payment, nil
}
