package TicketComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
)

// CreateTicketInfo .
func CreateTicketInfo(userTicket *UserTicketRepository.UserTicket) *ticket.TicketInfo {

	if userTicket == nil {
		return nil
	}

	return &ticket.TicketInfo{
		TicketId:    userTicket.TicketID,
		EventId:     userTicket.EventID,
		TicketName:  userTicket.Name,
		TicketDesc:  userTicket.Desc,
		TicketPrice: userTicket.Price,
	}

}
