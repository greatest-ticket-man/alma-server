package TicketRpcService

import (
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"context"
	"time"
)

// CreateTicket チケットの作成
func CreateTicket(ctx context.Context, mid string, txTime time.Time, eventID string, ticketID string, ticketName string,
	ticketPrice int32, ticketDesc string) bool {

	// TODO check

	// Create
	UserTicketRepository.Insert(ctx, txTime, ticketID, eventID, ticketName, ticketDesc, ticketPrice)
	return true
}
