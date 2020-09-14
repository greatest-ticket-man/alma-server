package TicketRpcService

import (
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/domain/ticket/TicketComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"context"
	"log"
	"time"
)

// UpdatePage チケットの編集画面
func UpdatePage(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string) *ticket.UpdatePageReply {

	// todo 権限
	userTicket := UserTicketRepository.FindOne(ctx, eventID, ticketID)

	log.Println("userTicket is ", jsonutil.Marshal(userTicket))

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
