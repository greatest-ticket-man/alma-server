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
			"mainTitle":   "ダッシュボード",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/content.html", "")),
			"script":      "",
			"css":         "",
		},
	)

}

// PageHTMLEmpty イベントが選択されていない場合はここにRedirect
func PageHTMLEmpty(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "ダッシュボード",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/empty/content.html", "")),
			"script":      "",
			"css":         template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/empty/css.html", "")), // 複数のファイルを指定させるためにhtmlを指定しています
		},
	)
}
