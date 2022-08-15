package service

import (
	"github.com/AndrewMislyuk/payments-api/pkg/stripe"
)

type Payments interface {
	ProductSubscription(productId string) (string, error)
}

type Service struct {
	Payments
}

func NewService(methodStripe *stripe.Stripe) *Service {
	return &Service{
		Payments: NewPaymentService(methodStripe),
	}
}