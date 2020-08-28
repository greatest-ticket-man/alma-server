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
