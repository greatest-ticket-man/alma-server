package TodoService

import (
	"alma-server/ap/src/repository/user/UserTodoRepository"
	"context"
	"time"
)

// CreateTodo todoを作成する
func CreateTodo(ctx context.Context, txTime time.Time, title string, desc string, name string) bool {

	todo := &UserTodoRepository.USER_TODO{
		Title:        title,
		Name:         name,
		Desc:         desc,
		CreateTime:   txTime,
		UpdateTime:   txTime,
		DeadlineTime: txTime.AddDate(0, 3, 0),
	}

	return UserTodoRepository.Insert(ctx, todo)
}

// GetTodoList todoのリストを取得する
func GetTodoList(ctx context.Context, name string) []*UserTodoRepository.USER_TODO {
	return UserTodoRepository.Find(ctx, name)
}
