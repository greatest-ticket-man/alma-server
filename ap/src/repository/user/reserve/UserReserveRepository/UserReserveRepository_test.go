package UserReserveRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/repository/user/reserve/UserReserveRepository"
	"context"
	"testing"
	"time"

	"github.com/franela/goblin"
	"go.mongodb.org/mongo-driver/bson"
)

// go test -timeout 30s alma-server/ap/src/repository/user/reserve/UserReserveRepository

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("UserReserveRepository:test", func() {
		ctx := context.Background()

		g.Before(func() {

			// 先に削除
			mongodb.GetUserCollection(ctx, UserReserveRepository.ThisCollectionName).DeleteOne(bson.M{UserReserveRepository.FReserveID: "test"})

		})

		// mid := "test"

		g.It("Add", func() {

			userReserve := &UserReserveRepository.UserReserve{
				ID:                "test",
				No:                1,
				EventID:           "test_event",
				CustomorID:        "test_customor",
				FirstName:         "テスト",
				FirstNameFurigana: "てすと",
				LastName:          "田中",
				LastNameFurigana:  "たなか",
				TicketID:          "test_ticketID",
				TicketCnt:         10,
				CreateTime:        time.Now(),
				UpdateTime:        time.Now(),
			}

			UserReserveRepository.Add(ctx, userReserve)

		})

	})

}
