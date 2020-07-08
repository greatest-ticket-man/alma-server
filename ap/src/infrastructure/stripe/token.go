package stripe

import (
	"alma-server/ap/src/common/error/chk"

	stripego "github.com/stripe/stripe-go/v71"
)

// token tokenの作成
// doc https://stripe.com/docs/api/tokens

// 基本ServerでTokenを作成することはないので、テストで使う用になると思われ

// CreateCardToken クレジットカードのTokenを作成する
func (c *AlmaStripe) CreateCardToken(cardNumber string, expMonth string, expYear string, cvc string) *stripego.Token {

	params := &stripego.TokenParams{
		Card: &stripego.CardParams{
			Number:   stripego.String(cardNumber),
			ExpMonth: stripego.String(expMonth),
			ExpYear:  stripego.String(expYear),
			CVC:      stripego.String(cvc),
		},
	}

	token, err := c.client.Tokens.New(params)
	chk.BE(err)

	return token
}
