package stripe_test

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/infrastructure/stripe"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/infrastructure/stripe

func Test(t *testing.T) {

	g := goblin.Goblin(t)
	test.Setup()

	g.Describe("stripe:test", func() {

		g.It("setup:test", func() {

			stripe.Setup(config.ConfigData.Stripe)

		})

	})

}
