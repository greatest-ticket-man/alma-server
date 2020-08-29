package EventRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/executor"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/domain/event/EventComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"alma-server/ap/src/repository/user/event/UserEventInviteMemberRepository"
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

// UpdatePage イベント編集ページ
func UpdatePage(ctx context.Context, mid string, eventID string) *event.UpdateEventPageReply {

	userEvent := UserEventRepository.Get(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	mstEventAuthList := MstEventAuthRepository.GetList()

	return &event.UpdateEventPageReply{
		EventName:         userEvent.Name,
		OrganizationName:  userEvent.Organization,
		EventAuthInfoList: EventComponent.CreateEventAuthInfoList(mstEventAuthList),
	}

}

// CreateEvent .
func CreateEvent(ctx context.Context, mid string, txTime time.Time, eventName string, organizationName string, inviteMemberList []*event.InviteMemberInfo) *event.CreateEventReply {

	eventID := uniqueidutil.GenerateUniqueID()

	// TODO 招待メンバーたちに、招待のMailを贈る
	userEventInviteMemberList := EventComponent.CreateInviteMemberList(eventID, txTime, inviteMemberList)

	var units []*executor.Unit

	// Event add
	units = append(units, UserEventRepository.CreateEventExecutor(ctx, txTime, eventID, eventName, organizationName))

	// Member add
	units = append(units, UserEventMemberRepository.CreateEventMemberExecutor(ctx, mid, txTime, eventID, "todo root"))

	// TempMemebrAdd .
	if len(userEventInviteMemberList) > 0 {
		units = append(units, UserEventInviteMemberRepository.BulkInsertInviteMemberExecutor(ctx, userEventInviteMemberList))
	}

	// execut
	executor.Do(units...)

	return &event.CreateEventReply{
		EventId:   eventID,
		EventName: eventName,
	}
}

// UpdateEvent .
func UpdateEvent(ctx context.Context, mid string, txTime time.Time, eventID string, eventName string, organizationName string, inviteMemberList []*event.InviteMemberInfo) bool {

	// TODO event を取得する

	// TODO 編集権限があるかを確認する

	// TODO
	userEvent := UserEventRepository.Get(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	// TODO errhandling 編集権限がありません

	userEventInviteMemberList := EventComponent.CreateInviteMemberList(eventID, txTime, inviteMemberList)

	var units []*executor.Unit
	// Event update

	units = append(units, UserEventRepository.UpdateEventExecutor(ctx, txTime, eventID, eventName, organizationName))

	// TempMemebrAdd .
	if len(userEventInviteMemberList) > 0 {
		units = append(units, UserEventInviteMemberRepository.BulkInsertInviteMemberExecutor(ctx, userEventInviteMemberList))
	}

	executor.Do(units...)

	return true
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

// GetEventList 自分が参加しているイベントリストを取得する
func GetEventList(ctx context.Context, mid string, txTime time.Time, searchText string) *event.GetEventListReply {

	// 自分が参加しているイベントを取得
	// この量は何万件になることは無いので、midでヒットするものをすべて取得する
	userEventMemberList := UserEventMemberRepository.GetList(ctx, mid)
	if len(userEventMemberList) == 0 {
		// なければそのまま返す
		return nil
	}

	eventIDList := EventComponent.CreateEventIDFromUserEventMember(userEventMemberList)
	userEventList := UserEventRepository.GetListInEventID(ctx, eventIDList)

	// TODO search textでfiltering

	return EventComponent.CreateGetEventListReply(userEventList)
}
