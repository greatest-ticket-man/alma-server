package stripe

import (
	"alma-server/ap/src/common/error/chk"

	stripego "github.com/stripe/stripe-go/v71"
)

// carge 支払い

// document https://stripe.com/docs/api/charges

// CreateCharge charge
// token is https://stripe.com/docs/payments/accept-a-payment-charges#web-create-token
func (c *AlmaStripe) CreateCharge(amount int64, currency stripego.Currency, desc string, token string) *stripego.Charge {

	params := &stripego.ChargeParams{
		Amount:      stripego.Int64(amount),
		Currency:    stripego.String(string(currency)),
		Description: stripego.String(desc),
		Source:      &stripego.SourceParams{Token: stripego.String(token)},
	}

	charge, err := c.client.Charges.New(params)
	chk.BE(err)
	return charge
}
