package stripe

import (
	stripeclient "github.com/stripe/stripe-go/v71/client"
)

// AlmaStripe StripeのWrapper
type AlmaStripe struct {
	client *stripeclient.API
}

// GetAllConsumer 全ての顧客をリストする
func (c *AlmaStripe) GetAllConsumer() {

}
