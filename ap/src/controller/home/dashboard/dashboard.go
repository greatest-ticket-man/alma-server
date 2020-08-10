package dashboard

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// PageHTML Dashboard
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "Dashboard",
			"mainContent": "ダッシュボードコンテンツ",
			"script":      "",
			"css":         "",
		},
	)

}
