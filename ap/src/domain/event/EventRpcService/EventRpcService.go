package EventRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/executor"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/domain/event/EventComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"alma-server/ap/src/repository/user/event/UserEventMemberRepository"
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"context"
	"time"
)

// CreatePage イベント作成ページ
func CreatePage() *event.CreateEventPageReply {

	mstEventAuthList := MstEventAuthRepository.GetList()

	return &event.CreateEventPageReply{
		EventAuthInfoList: EventComponent.CreateEventAuthInfoList(mstEventAuthList),
	}
}

// CreateEvent .
// TODO memberInfoをinviteMemberInfoにする
func CreateEvent(ctx context.Context, mid string, txTime time.Time, eventName string, organizationName string, memberList []*event.MemberInfo) *event.CreateEventReply {

	eventID := uniqueidutil.GenerateUniqueID()

	// TODO 招待メンバーたちに、招待のMailを贈る

	var units []*executor.Unit

	// Event add
	units = append(units, UserEventRepository.CreateEventExecutor(ctx, txTime, eventID, eventName, organizationName))

	// Member add
	// 自分をRootユーザーで登録する
	units = append(units, UserEventMemberRepository.CreateEventMemberExecutor(ctx, mid, txTime, eventID, "todo root"))

	// TempMemebrAdd .
	// units = append(units, UserEventInviteMemberRepository.BulkInsertInviteMemberExecutor(ctx))

	// execut
	executor.Do(units...)

	return &event.CreateEventReply{
		EventId:   eventID,
		EventName: eventName,
	}
}

// UpdateEvent .
func UpdateEvent(ctx context.Context, mid string, txTime time.Time, eventID string, eventName string, organizationName string, memberList []*event.MemberInfo) bool {

	// TODO event を取得する

	// TODO 編集権限があるかを確認する

	// TODO
	userEvent := UserEventRepository.Get(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	// TODO errhandling 編集権限がありません

	// tempMemberInfoList := EventComponent.CreateTempMemberInfoList(txTime, memberList)

	// update
	UserEventRepository.Update(ctx, txTime, eventID, eventName, organizationName)

	return false
}

// GetEvent イベントのデータを取得する
func GetEvent(ctx context.Context, mid string, eventID string) *event.HomeReply {

	// TODO 下記のEventにアクセスする権限があるかを確認する

	userEvent := UserEventRepository.Get(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	return &event.HomeReply{
		EventId:   eventID,
		EventName: userEvent.Name,
	}
}
