package stripe

import (
	"fmt"
	"time"
)

type Stripe struct {
}

func NewStripe() *Stripe {
	return &Stripe{}
}

func (s *Stripe) GetPaymentURL(productId string, c chan string) {
	time.Sleep(time.Second * 3)

	url := fmt.Sprintf("https://api.stripe.com/subscribe/%s", productId)

	c <- url
}
