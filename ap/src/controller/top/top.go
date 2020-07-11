package top

import (
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/CommonHTMLService"
	"net/http"
)

// PageHTML top画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/top/page.html",
		map[string]interface{}{
			"head":   CommonHTMLService.GetHead(),
			"header": CommonHTMLService.GetHeader(),
			"footer": CommonHTMLService.GetFooter(),
		},
	)

}
