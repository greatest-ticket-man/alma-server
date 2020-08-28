package UserEventRepository

import (
	"alma-server/ap/src/common/executor"
	"context"
	"time"
)

// CreateEventExecutor イベントを作成するExecutor
func CreateEventExecutor(ctx context.Context, txTime time.Time, eventID string, name string, organization string) *executor.Unit {

	u := &executor.Unit{}

	u.Execute = func() interface{} {
		u.ExecuteResult = Insert(ctx, txTime, eventID, name, organization)
		return u.ExecuteResult
	}

	u.Rollback = func(ctx context.Context) interface{} {
		u.RollbackResult = Remove(ctx, eventID)
		return u.RollbackResult
	}

	return u
}

// UpdateEventExecutor イベント情報を編集するExecutor
func UpdateEventExecutor(ctx context.Context, txTime time.Time, eventID string, name string, organization string) *executor.Unit {

	u := &executor.Unit{}

	u.Execute = func() interface{} {
		u.ExecuteResult = FindOneAndUpdate(ctx, txTime, eventID, name, organization)
		return u.ExecuteResult
	}

	u.Rollback = func(ctx context.Context) interface{} {

		before := u.ExecuteResult.(*UserEvent)

		if before == nil {
			u.RollbackResult = Remove(ctx, eventID)
		} else {
			u.RollbackResult = Update(ctx, txTime, eventID, before.Name, before.Organization)
		}

		return u.RollbackResult
	}

	return u
}
