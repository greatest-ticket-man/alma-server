package create

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"html/template"
	"net/http"
)

// PageHTML イベント作成画面
func PageHTML(w http.ResponseWriter, r *http.Request) {
	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "イベントの作成",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/create/content.html", "")),
			"script":      template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/create/script.html", "")),
			"css":         template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/create/css.html", "")),
		},
	)
}
