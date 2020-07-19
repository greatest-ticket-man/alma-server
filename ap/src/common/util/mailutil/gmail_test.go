package mailutil_test

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/mailutil"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/util/mailutil

func Test(t *testing.T) {

	g := goblin.Goblin(t)
	test.Setup()

	g.Describe("gmail:test", func() {

		g.It("Gmailが送信できていることを確認する", func() {
			err := mailutil.SendGmail("sunjin110@gmail.com", "Golangからのテスト送信", "今日はいい天気ですね")
			chk.SE(err)
		})

	})

}
