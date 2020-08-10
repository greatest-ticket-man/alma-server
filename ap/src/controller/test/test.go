package test

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/response"
	"log"
	"net/http"
)

// PageHTML UIテスト画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	mid := almactx.GetMid(r.Context())

	log.Println("mid is ", mid)

	response.HTML(
		w,
		"/template/controller/test/page.html",
		map[string]interface{}{},
	)
}
