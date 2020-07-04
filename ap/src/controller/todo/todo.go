package todo

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/TodoService"
	"encoding/json"
	"html/template"
	"net/http"
	"time"
)

// PageHTML todoの追加画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	head := htmlutil.CreateTemplateToString("/template/common/head.html", "")
	header := htmlutil.CreateTemplateToString("/template/common/header.html", "")
	footer := htmlutil.CreateTemplateToString("/template/common/footer.html", "")

	// TODO list

	response.HTML(
		w,
		"/template/controller/todo/page.html",
		map[string]interface{}{
			"head":   template.HTML(head),
			"header": template.HTML(header),
			"footer": template.HTML(footer),
		},
	)
}

// CreateTodo todoの作成
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	data := &struct {
		Title string `json:"title"`
		Desc  string `json:"desc"`
	}{}
	err := json.NewDecoder(r.Body).Decode(data)
	chk.SE(err)

	// CreateTodo
	TodoService.CreateTodo(r.Context(), time.Now(), data.Title, data.Desc, "sunjin")

	response.JSON(w, true)
}
