package TicketService

import (
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"context"
)

// GetUserTicketMap .
func GetUserTicketMap(ctx context.Context, eventID string) map[string]*UserTicketRepository.UserTicket {

	ticketMap := map[string]*UserTicketRepository.UserTicket{}

	userTicketList := UserTicketRepository.Find(ctx, eventID)
	for _, userTicket := range userTicketList {
		ticketMap[userTicket.TicketID] = userTicket
	}

	return ticketMap
}
