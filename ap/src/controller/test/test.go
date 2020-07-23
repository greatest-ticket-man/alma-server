package test

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// PageHTML UIテスト画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/test.page.html",
		map[string]interface{}{},
	)
}
