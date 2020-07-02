package config_test

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/util/jsonutil"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/config

func Test(t *testing.T) {

	g := goblin.Goblin(t)

	g.Describe("config: test", func() {

		g.It("configが正しく受け取れていることを確認する", func() {

			log.Println("hello")

			config := config.Setup("../../../config/local.toml")
			log.Println("config is ", jsonutil.Unmarshal(config))

			g.Assert(config.MongoDatabases != nil).IsTrue("mongoの設定がありません")
			g.Assert(config.HTTPServer != nil).IsTrue("httpserverの設定がありません")

		})

	})

}
