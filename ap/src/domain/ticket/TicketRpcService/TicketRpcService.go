package TicketRpcService

import (
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/ticket/TicketComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"context"
	"time"
)

// Page チケットのトップページ画面
func Page(ctx context.Context, mid string, txTime time.Time, eventID string) *ticket.PageReply {

	userTicketList := UserTicketRepository.Find(ctx, eventID)
	eventName := EventService.GetEventName(ctx, eventID)

	return &ticket.PageReply{
		TicketInfoList: TicketComponent.CreateTicketInfoList(userTicketList),
		EventName:      eventName,
		EventId:        eventID,
	}
}

// UpdatePage チケットの編集画面
func UpdatePage(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string) *ticket.UpdatePageReply {

	// todo 権限
	userTicket := UserTicketRepository.FindOne(ctx, eventID, ticketID)

	eventName := EventService.GetEventName(ctx, eventID)

	return &ticket.UpdatePageReply{
		EventName:  eventName,
		TicketInfo: TicketComponent.CreateTicketInfo(userTicket),
	}
}

// CreateTicket チケットの作成
func CreateTicket(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string, name string,
	price int32, desc string, scheduleStockInfoList []*ticket.TicketScheduleStockInfo) bool {

	// TODO check
	// TODO EventIDが指定しているものと正しいか

	// Create
	userTicket := TicketComponent.CreateUserTicket(txTime, eventID, ticketID, name, price, desc, scheduleStockInfoList)
	UserTicketRepository.Insert(ctx, userTicket)

	return true
}

// UpdateTicket チケットの編集
func UpdateTicket(ctx context.Context, mid string, txTime time.Time, eventID string, ticketInfo *ticket.TicketInfo) bool {

	// make scheduleStockInfoMap
	scheduleStockMap := TicketComponent.CreasteScheduleStockMap(ticketInfo.ScheduleStockList)

	// update
	UserTicketRepository.Update(ctx, txTime, ticketInfo.TicketId, eventID, ticketInfo.Name, ticketInfo.Desc, ticketInfo.Price, scheduleStockMap)

	return true
}

// DeleteTicket チケットの削除
func DeleteTicket(ctx context.Context, mid string, txTime time.Time, eventID string, ticketIDList []string) int32 {

	// TODO Action Log

	return UserTicketRepository.RemoveMany(ctx, eventID, ticketIDList)
}
