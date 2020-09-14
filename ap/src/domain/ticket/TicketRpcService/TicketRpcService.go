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
func CreateTicket(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string, ticketName string,
	ticketPrice int32, ticketDesc string) bool {

	// TODO check

	// Create
	UserTicketRepository.Insert(ctx, txTime, ticketID, eventID, ticketName, ticketDesc, ticketPrice)
	return true
}
