package UserEventInviteMemberRepository

import (
	"alma-server/ap/src/common/executor"
	"context"
)

// BulkInsertInviteMemberExecutor 招待データを一気に追加する
func BulkInsertInviteMemberExecutor(ctx context.Context, userEventInviteMemberList []*UserEventInviteMember) *executor.Unit {

	u := &executor.Unit{}

	u.Execute = func() interface{} {
		u.ExecuteResult = InsertBulk(ctx, userEventInviteMemberList)
		return u.ExecuteResult
	}

	u.Rollback = func(ctx context.Context) interface{} {

		insertedIDList := u.ExecuteResult.([]interface{})

		var emailList []string
		for _, insertedID := range insertedIDList {
			email := insertedID.(string)
			emailList = append(emailList, email)
		}

		u.RollbackResult = RemoveMany(ctx, emailList)
		return u.RollbackResult
	}

	return u
}
