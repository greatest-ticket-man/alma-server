package MstEventAuthRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/master/authority/MstEventAuthRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MstEventAuthRepository:test", func() {

		g.It("データが受け取れることを確認する", func() {

			result := MstEventAuthRepository.Get("test-data-1")

			g.Assert(result == nil).IsFalse("データが取得できませんでした")

		})

		g.It("リストを受け取る", func() {

			list := MstEventAuthRepository.GetList()

			g.Assert(len(list) == 0).IsFalse()

		})

	})

}
