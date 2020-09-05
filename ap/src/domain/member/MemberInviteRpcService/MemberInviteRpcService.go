package MemberInviteRpcService

import (
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/member/MemberComponent"
	memberinvite "alma-server/ap/src/infrastructure/grpc/proto/member/invite"
	"alma-server/ap/src/repository/user/event/UserEventInviteMemberRepository"
	"context"
	"time"
)

// PageHTML .
func PageHTML(ctx context.Context, mid string, txTime time.Time, eventID string) *memberinvite.PageHTMLReply {

	userEvent := EventService.GetEvent(ctx, eventID)

	userEventInviteMemberList := UserEventInviteMemberRepository.GetList(ctx, eventID)

	return &memberinvite.PageHTMLReply{
		EventName:            userEvent.Name,
		MemberInviteInfoList: MemberComponent.CreateMemberInviteInfoList(userEventInviteMemberList),
	}

}
