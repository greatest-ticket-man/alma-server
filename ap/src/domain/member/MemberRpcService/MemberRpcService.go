package MemberRpcService

import (
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/member/MemberComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/member"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"alma-server/ap/src/repository/user/event/UserEventMemberRepository"
	"context"
	"time"
)

// PageHTML .
func PageHTML(ctx context.Context, mid string, txTime time.Time, eventID string) *member.PageHTMLReply {

	userEvent := EventService.GetEvent(ctx, eventID)

	userEventMemberList := UserEventMemberRepository.GetListFromEventID(ctx, eventID)

	memberMidList := MemberComponent.GetMidListFromUserEventMember(userEventMemberList)

	userAccountList := UserAccountRepository.GetList(ctx, memberMidList)

	return &member.PageHTMLReply{
		EventName:      userEvent.Name,
		MemberInfoList: MemberComponent.CreateMemberInfoList(userEventMemberList, userAccountList),
	}
}
