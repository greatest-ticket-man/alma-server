package MstTicketPayTypeRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/repository/master/ticket/MstTicketPayTypeRepository"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/master/ticket/MstTicketPayTypeRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MstPayTypeRepository:test", func() {

		g.It("データが受け取れることを確認する", func() {

			result := MstTicketPayTypeRepository.GetMap()
			g.Assert(len(result) == 0).IsFalse()

		})

	})

}
