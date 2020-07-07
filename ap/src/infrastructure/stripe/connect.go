package stripe

import (
	"alma-server/ap/src/common/config"

	stripeclient "github.com/stripe/stripe-go/v71/client"
)

var client *stripeclient.API

// Setup Stripe Connect
func Setup(config *config.Stripe) bool {

	sc := &stripeclient.API{}
	sc.Init(config.SecretKey, nil)
	client = sc

	return true
}

// GetClient clientを取得する
func GetClient() *stripeclient.API {
	return client
}
