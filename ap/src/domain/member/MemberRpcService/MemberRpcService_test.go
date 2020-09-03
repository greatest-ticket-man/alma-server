package MemberRpcService_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/member/MemberRpcService"
	"context"
	"log"
	"testing"
	"time"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/domain/member/MemberRpcService

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MemberRpcService:test", func() {

		ctx := context.Background()
		mid := "bsof9voul2pukksa2v3g"
		txTime := time.Now()

		g.It("PageHTML", func() {
			// eventIDを適当に取得
			reply := EventRpcService.CreateEvent(ctx, mid, txTime, "aa", "a", nil)

			result := MemberRpcService.PageHTML(ctx, mid, txTime, reply.EventId)

			log.Println("result is ", jsonutil.Marshal(result))

		})

	})

}
