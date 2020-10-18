package ReserveRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/reserve/ReserveComponent"
	"alma-server/ap/src/domain/ticket/TicketComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
	"alma-server/ap/src/repository/user/reserve/UserReserveRepository"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"context"
	"time"
)

// Page 予約ページのトップを取得する
func Page(ctx context.Context, mid string, txTime time.Time, eventID string) *reserve.PageReply {

	userEvent := EventService.GetEvent(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	// reserve
	userReserveList := UserReserveRepository.GetList(ctx, eventID)

	reserveInfoList := ReserveComponent.CreateReserveInfoList(userReserveList)

	return &reserve.PageReply{
		EventId:         eventID,
		EventName:       userEvent.Name,
		ReserveInfoList: reserveInfoList,
	}
}

// CreatePage 予約作成画面を取得する
func CreatePage(ctx context.Context, mid string, txTime time.Time, eventID string) *reserve.CreatePageReply {

	userEvent := EventService.GetEvent(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	// GetTicketAll
	userTicketList := UserTicketRepository.Find(ctx, eventID)

	return &reserve.CreatePageReply{
		EventId:        eventID,
		EventName:      userEvent.Name,
		TicketInfoList: TicketComponent.CreateTicketInfoList(userTicketList),
	}
}
