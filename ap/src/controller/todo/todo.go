package todo

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
	"html/template"
)

// PageHTML todoの追加画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	head := htmlutil.CreateTemplateToString("/template/common/head.html", "")
	header := htmlutil.CreateTemplateToString("/template/common/header.html", "")
	footer := htmlutil.CreateTemplateToString("/template/common/footer.html", "")


	response.HTML(
		w,
		"/template/controller/todo/page.html",
		map[string]interface{}{
			"head": template.HTML(head),
			"header": template.HTML(header),
			"footer": template.HTML(footer),
		},
	)
}
