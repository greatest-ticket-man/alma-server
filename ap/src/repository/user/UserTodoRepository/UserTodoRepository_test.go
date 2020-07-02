package UserTodoRepository_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/repository/user/UserTodoRepository"
	"context"
	"log"
	"testing"
	"time"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/repository/user/UserTodoRepository

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	test.Setup()

	g.Describe("UserTodoRepository:test", func() {

		g.It("Insertできるかのテスト", func() {

			userTodo := &UserTodoRepository.USER_TODO{
				Name:         "sunjin",
				Title:        "実装",
				Desc:         "チケットのクレジット決済機能を作成する",
				DeadlineTime: time.Date(2020, 8, 1, 12, 0, 0, 0, time.UTC),
				CreateTime:   time.Now(),
				UpdateTime:   time.Now(),
			}

			result := UserTodoRepository.Insert(context.Background(), userTodo)
			g.Assert(result).IsTrue("追加に失敗しました")
		})

		g.It("Findできること", func() {

			userTodoList := UserTodoRepository.Find(context.Background(), "sunjin")

			log.Println("result is ", jsonutil.Marshal(userTodoList))

			g.Assert(len(userTodoList) == 0).IsFalse("データが見つかりませんでした")

		})

	})

}
