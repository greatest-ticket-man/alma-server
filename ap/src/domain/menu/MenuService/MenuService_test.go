package MenuService_test

import (
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/domain/menu/MenuService"
	"log"
	"testing"

	"github.com/franela/goblin"
)

// go test -v -cover=1 -timeout 30s alma-server/ap/src/domain/menu/MenuService


	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("MenuService:test", func() {

		g.It("GetMenu", func() {

			result := MenuService.GetMenu("test")
			log.Println("result is ", jsonutil.Marshal(result))

		})

	})

}
