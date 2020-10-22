package UserSequenceRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/repository/user/UserSequenceRepository"
	"context"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/user/UserSquenceRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("UserSequenceRepository:test", func() {

		g.It("IncTest:test", func() {

			UserSequenceRepository.Next(context.Background(), "test_eid", UserSequenceRepository.ReserveKey)

		})

	})
}
