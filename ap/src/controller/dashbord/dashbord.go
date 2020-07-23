package dashbord

import (
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/CommonHTMLService"
	"net/http"
)

// PageHTML dashbord画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/dashbord/page.html",
		map[string]interface{}{
			"head":   CommonHTMLService.GetHead(),
			"header": CommonHTMLService.GetHeader(),
			"footer": CommonHTMLService.GetFooter(),
		},
	)

}
