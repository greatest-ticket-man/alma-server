package jwt_test

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/test"
	"alma-server/ap/src/common/util/jsonutil"
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

			token := jwt.Auth(req)
			log.Println("token is ", jsonutil.Marshal(token))

		})

		g.It("実際とは違うTokenを送る,パニックが発生すること", func() {

			req, err := http.NewRequest("POST", "/hhhh", nil)
			chk.SE(err)

			// header
			req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTcwODA5MTQsIm5hbWUiOiJ0ZXN0QHRlc3QuY29tIn0.aVoDUiOeRyqTLd-yMpq86BG0cDhwAudACaV8Z59eGl-KpRtS9s6n_N3ovsqoHxLyVQiM4LZIKJsWn1Hz9GbC2uWTqvhYC9EZ0e2_k9MaJ8ZhX4fBqFMLFxOvFEnOaxyKLa8EAu5O1oocvgAPjE1VgWQJu5HhgcTucsH9DTgxNSrHGOOscu3ue5yRJFzu0_jNQdbbeD4OvLxj7aN5mCrb-d6p3zt4B-8hYpHKk7yjEVQxvV-2Y0j2Z9i5OwgpQ3AgHSqJOlz-069lLAKlmEMtd4VtrMt9hbsYbjjh2zJfm-hzOdL8IzjylEwKras4HankqYLMkUhjQiaw7xJZ_iQx1agZ_3oumGclnU0cLAQMHK-SLzUg6vc8NHEi2w3PBlrrVcB4LK2WTKi5UetLMJToyh69EnhK4KUyCYrt3LiSeUkDSNHvA0LASjuuTXwu5VRk7BA7K4JLFClmmWuWidDGnwBHaXl23235KaEuekO64BmFxfFzSE-2gUAwUDCeawg-"))

			defer func() {
				if err := recover(); err != nil {
					log.Println("ok")
				}
			}()

			_ = jwt.Auth(req)

			g.Fail("エラーが発生しませんでした")
		})

		g.It("有効時間外のTokenを送信したとき", func() {

			txTime := time.Now().Add(-24 * time.Hour)

			tokenStr := jwt.New(txTime, "sunjin@sunjin.com", "sunjin")

			req, err := http.NewRequest("GET", "/aaa", nil)
			chk.SE(err)

			// header
			req.Header.Set("Authorization", fmt.Sprintf("BEARER %s", tokenStr))

			defer func() {
				if err := recover(); err != nil {
					if fmt.Sprintf("%v", err) != "Token is expired" {
						g.Fail("想定していないエラー")
					}
				}
			}()

			t := jwt.Auth(req)

			g.Fail("エラーが発生しませんでした")
			log.Println("t is ", jsonutil.Marshal(t))

		})

	})

}
