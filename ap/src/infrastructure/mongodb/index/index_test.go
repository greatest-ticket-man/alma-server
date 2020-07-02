package index_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/infrastructure/mongodb/index"
	"context"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/infrastructure/mongodb/index

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("index:test", func() {

		test.Setup()

		g.It("indexが追加されるかを確認する", func() {

			index.CreateIndex(context.Background())

		})

	})

}
