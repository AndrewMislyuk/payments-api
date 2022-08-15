package service

import (
	"errors"

	"github.com/AndrewMislyuk/payments-api/pkg/stripe"
)

type Payment struct {
	methodStripe *stripe.Stripe
}

func NewPaymentService(str *stripe.Stripe) *Payment {
	return &Payment{
		methodStripe: str,
	}
}

func (p *Payment) ProductSubscription(productId string) (string, error) {
	if productId == "" {
		return "", errors.New("url mustn't be empty")
	}
	
	c := make(chan string)
	go p.methodStripe.GetPaymentURL(c)
	url := <-c

	return url, nil
}
