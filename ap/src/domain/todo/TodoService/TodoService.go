package TodoService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/repository/user/UserTodoRepository"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

// RemoveTodo todoを削除する
func RemoveTodo(ctx context.Context, id *primitive.ObjectID) bool {

	result := UserTodoRepository.Delete(ctx, id)
	if result != 1 {
		// todoの削除に失敗しました
		chk.LE(errmsg.TodoDeleteFailed)
	}

	return true
}

// GetTodoList todoのリストを取得する
func GetTodoList(ctx context.Context, name string) []*UserTodoRepository.USER_TODO {
	return UserTodoRepository.Find(ctx, name)
}
