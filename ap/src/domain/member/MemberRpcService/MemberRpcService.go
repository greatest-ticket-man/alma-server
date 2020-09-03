package MemberRpcService

import (
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"context"
	"time"
)

// PageHTML .
func PageHTML(ctx context.Context, mid string, txTime time.Time, eventID string) *common.EventNameReply {

	userEvent := EventService.GetEvent(ctx, eventID)

	// TODO get member page

	// TODO paging

	return &common.EventNameReply{
		EventName: userEvent.Name,
	}
}
