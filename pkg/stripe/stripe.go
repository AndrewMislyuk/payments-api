package stripe

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Stripe struct {
}

func NewStripe() *Stripe {
	return &Stripe{}
}

func (s *Stripe) GetPaymentURL(c chan string) {
	time.Sleep(time.Second * 3)

	url := fmt.Sprintf("https://api.stripe.com/subscribe/%s", uuid.New().String())

	c <- url
}
