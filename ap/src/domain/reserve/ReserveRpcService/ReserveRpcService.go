package ReserveRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/reserve/ReserveComponent"
	"alma-server/ap/src/domain/sequence/SequenceService"
	"alma-server/ap/src/domain/ticket/TicketComponent"
	"alma-server/ap/src/domain/ticket/TicketService"
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
	"alma-server/ap/src/repository/master/ticket/MstTicketPayTypeRepository"
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

	// ticketMap
	userTicketMap := TicketService.GetUserTicketMap(ctx, eventID)

	// ticketPayTypeMap
	mstTicketPayTypeMap := MstTicketPayTypeRepository.GetMap()

	reserveInfoList := ReserveComponent.CreateReserveInfoList(userReserveList, userTicketMap, mstTicketPayTypeMap)

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

	// ticket pay type
	mstTicketPayTypeList := MstTicketPayTypeRepository.GetList()
	ticketPayTypeList := TicketComponent.CreateTicketPayTypeList(mstTicketPayTypeList)

	return &reserve.CreatePageReply{
		EventId:           eventID,
		EventName:         userEvent.Name,
		TicketInfoList:    TicketComponent.CreateTicketInfoList(userTicketList),
		TicketPayTypeList: ticketPayTypeList,
	}
}

// CreateReserve 予約を作成する
func CreateReserve(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string, ticketNum int32, desc string, name string, nameFrigana string, email string, payTypeID string) bool {
	userEvent := EventService.GetEvent(ctx, eventID)
	if userEvent == nil {
		// イベントが存在しません
		chk.LE(errmsg.EventNotFound)
	}

	// TODO stockを確認

	// TODO customerに追加する

	// check
	mstTicketPayType := MstTicketPayTypeRepository.Get(payTypeID)
	if mstTicketPayType == nil {
		// 存在しない販売形式が指定されました
		chk.LE(errmsg.TicketNonExistSalesFormat)
	}

	// counter
	seq := SequenceService.NextReserveSeq(ctx, eventID)

	// 追加
	reserveID := uniqueidutil.GenerateUniqueID()

	UserReserveRepository.Insert(ctx, txTime, reserveID, seq, eventID, "", name, nameFrigana, email, ticketID, ticketNum, payTypeID)
	return true
}
