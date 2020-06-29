package login

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"html/template"
	"net/http"
)

// PageHTML ログイン画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	head := htmlutil.CreateTemplateToString("/template/common/head/head.html", "")

	response.HTML(
		w,
		"/template/controller/login/page.html",
		map[string]interface{}{
			"head": template.HTML(head), // template.HTML型に変換すれば、エスケープされなくなる
		},
	)
}
