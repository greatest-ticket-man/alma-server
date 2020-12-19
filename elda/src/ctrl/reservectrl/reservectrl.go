package reservectrl

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// PageHTML .
func PageHTML(w http.ResponseWriter, r *http.Request) {
	// 予約画面
	response.HTML(w, "/template/ctrl/reserve/reserve.html", map[string]interface{}{})
}
