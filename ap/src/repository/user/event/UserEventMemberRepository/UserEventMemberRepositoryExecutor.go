package UserEventMemberRepository

import (
	"alma-server/ap/src/common/executor"
	"context"
	"time"
)

// CreateEventMemberExecutor メンバーを一人追加するExecutor
func CreateEventMemberExecutor(ctx context.Context, mid string, txTime time.Time, eventID string, authID string) *executor.Unit {

	u := &executor.Unit{}

	u.Execute = func() interface{} {
		u.ExecuteResult = FindOneAndUpsert(ctx, mid, txTime, eventID, authID)
		return u.ExecuteResult
	}

	u.Rollback = func(ctx context.Context) interface{} {

		if u.ExecuteResult == nil {
			u.RollbackResult = Remove(ctx, mid, txTime, eventID)
		} else {
			before := u.ExecuteResult.(*UserEventMember)
			u.RollbackResult = Upsert(ctx, mid, txTime, eventID, before.AuthID)
		}

		return u.RollbackResult
	}

	return u
}
