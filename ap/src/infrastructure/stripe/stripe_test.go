package stripe_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/infrastructure/stripe"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/infrastructure/stripe

func Test(t *testing.T) {

	g := goblin.Goblin(t)
	test.Setup()

	g.Describe("stripe:test", func() {

		almaStripe := stripe.GetClient()

		g.It("GetAllProduct", func() {
			result := almaStripe.GetAllProductList()
			log.Println("result is ", jsonutil.Marshal(result))
			g.Assert(len(result) == 0).IsFalse("商品が取得できていません")
		})

	})

}
