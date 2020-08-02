package UserAccountRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"log"
	"testing"
	"time"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/user/UserAccountRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("UserAccountRepository:Test", func() {

		ctx := context.Background()
		txTime := time.Now()

		g.It("Insert:Test", func() {

			mid := uniqueidutil.GenerateUniqueID()

			// ObjectIDが取得できること
			result := UserAccountRepository.Insert(ctx, txTime, mid, "test@test.com", "test password")

			log.Println("result is ", result)
		})

	})

}
