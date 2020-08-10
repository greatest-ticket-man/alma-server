package dashboard

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"html/template"
	"net/http"
)

// PageHTML Dashboard
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "Dashboard",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/component/home/dashboard/content.html", "")),
			"script":      "",
			"css":         "",
		},
	)

}
