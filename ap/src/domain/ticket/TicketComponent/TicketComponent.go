package TicketComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
)

// CreateTicketInfoList .
func CreateTicketInfoList(userTicketList []*UserTicketRepository.UserTicket) []*ticket.TicketInfo {
	var ticketInfoList []*ticket.TicketInfo
	for _, userTicket := range userTicketList {
		ticketInfoList = append(ticketInfoList, CreateTicketInfo(userTicket))
	}

	return ticketInfoList
}

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
