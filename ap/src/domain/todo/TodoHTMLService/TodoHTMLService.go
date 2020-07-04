package TodoHTMLService

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/domain/todo/TodoService"
	"context"
	"html/template"
)

// GetTodoListTable todoのtableを作成する
func GetTodoListTable(ctx context.Context, name string) template.HTML {

	todoList := TodoService.GetTodoList(ctx, name)

	return template.HTML(htmlutil.CreateTemplateToString("/template/controller/todo/table.html", map[string]interface{}{
		"todoList": todoList,
	}))
}
