package todo

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/CommonHTMLService"
	"alma-server/ap/src/domain/todo/TodoHTMLService"
	"alma-server/ap/src/domain/todo/TodoService"
	"encoding/json"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PageHTML todoの追加画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/todo/page.html",
		map[string]interface{}{
			"head":   CommonHTMLService.GetHead(),
			"header": CommonHTMLService.GetHeader(),
			"footer": CommonHTMLService.GetFooter(),
			"table":  TodoHTMLService.GetTodoListTable(r.Context(), "sunjin"),
			"form":   TodoHTMLService.GetTodoForm(r.Context()),
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

// RemoveTodo todoの削除
func RemoveTodo(w http.ResponseWriter, r *http.Request) {

	data := &struct {
		ID *primitive.ObjectID `json:"id"`
	}{}

	err := json.NewDecoder(r.Body).Decode(data)
	chk.SE(err)

	TodoService.RemoveTodo(r.Context(), data.ID)

	response.JSON(w, true)
}
