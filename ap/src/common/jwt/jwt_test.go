package jwt_test

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/test"
	"fmt"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/franela/goblin"
)

// go test -v -count=1 -timeout 30s alma-server/ap/src/common/jwt

func Test(t *testing.T) {

	test.Setup()

	g := goblin.Goblin(t)

	g.Describe("Jwt:Test", func() {

		g.It("New", func() {
			log.Println("Jwt Test")

			tokenStr := jwt.New(time.Now(), "test@test.com", "test")
			log.Println("tokenStr is ", tokenStr)

			g.Assert(tokenStr == "").IsFalse()
		})

		g.It("認証", func() {

			tokenStr := jwt.New(time.Now(), "test@test.com", "test")

			req, err := http.NewRequest("GET", "/hoge", nil)
			chk.SE(err)

			// header
			req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", tokenStr))

			jwt.Auth(req)

		})

	})

}
