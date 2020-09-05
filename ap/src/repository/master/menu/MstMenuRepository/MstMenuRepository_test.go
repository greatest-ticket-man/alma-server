package MstMenuRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/repository/master/menu/MstMenuRepository"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/master/menu/MstMenuRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MstMenuRepository:test", func() {

		g.It("データが受け取れることを確認する", func() {

			result := MstMenuRepository.GetMap()

			g.Assert(len(result) == 0).IsFalse()

			log.Println("result is ", jsonutil.Marshal(result))

		})

	})

}
