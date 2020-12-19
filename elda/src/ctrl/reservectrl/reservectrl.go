package reservectrl

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// PageHTML .
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// TODO EventIDがない場合は、Error

	// TODO そのイベントの情報公開日より前の場合はError

	// TODO 大丈夫であれば、チケットとかの情報を表示してあげる
	// 予約画面
	response.HTML(w, "/template/ctrl/reserve/reserve.html", map[string]interface{}{})
}
