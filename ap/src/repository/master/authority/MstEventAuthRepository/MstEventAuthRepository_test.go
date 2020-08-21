package MstEventAuthRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/master/authority/MstEventAuthRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MstEventAuthRepository:test", func() {

		g.It("データが受け取れることを確認する", func() {

			result := MstEventAuthRepository.Get("test-data")

			log.Println("result is ", jsonutil.Marshal(result))

			g.Assert(result == nil).IsFalse("データが取得できませんでした")

		})

	})

}
