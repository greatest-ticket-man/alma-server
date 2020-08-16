package UserAccountRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"testing"
	"time"

	"github.com/franela/goblin"
	"gopkg.in/mgo.v2/bson"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/user/UserAccountRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("UserAccountRepository:Test", func() {

		ctx := context.Background()
		txTime := time.Now()
		email := "test@test.com"

		g.It("Insert:Test", func() {

			// 先に削除
			mongodb.GetUserCollection(ctx, UserAccountRepository.ThisCollectionName).DeleteOne(bson.M{
				UserAccountRepository.FEmail: email,
			})

			mid := uniqueidutil.GenerateUniqueID()

			// ObjectIDが取得できること
			result := UserAccountRepository.Insert(ctx, txTime, mid, "test_name", "test@test.com", "test password")

			g.Assert(mid).Eql(result)
		})

		g.It("データを取得できていることを確認", func() {

			result := UserAccountRepository.GetFromEmail(ctx, email)

			g.Assert(result.Email).Eql(email)

		})

	})

}
