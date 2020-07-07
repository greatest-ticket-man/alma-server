package httputil_test

import (
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/util/httputil

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("httputil", func() {

		g.It("test", func() {

			log.Println("hoge")

		})

	})

}
