package stripe

import (
	"alma-server/ap/src/common/config"
	"log"

	client "github.com/stripe/stripe-go/v71/client"
)

// Setup Stripe Connect
func Setup(config *config.Stripe) bool {

	sc := &client.API{}
	sc.Init(config.SecretKey, nil) //

	// https://github.com/stripe/stripe-go

	log.Println("sc is ", sc)

	return true

}
