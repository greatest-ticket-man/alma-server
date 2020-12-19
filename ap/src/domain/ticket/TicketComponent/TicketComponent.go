package TicketComponent

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
	"alma-server/ap/src/repository/master/ticket/MstTicketPayTypeRepository"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
	"time"
)

// CreateTicketInfoList .
func CreateTicketInfoList(userTicketList []*UserTicketRepository.UserTicket) []*ticket.TicketInfo {
	var ticketInfoList []*ticket.TicketInfo
	for _, userTicket := range userTicketList {
		ticketInfoList = append(ticketInfoList, CreateTicketInfo(userTicket))
	}

	return ticketInfoList
}

// CreateUserTicket .
// MapのIDは新しく作成します
func CreateUserTicket(txTime time.Time, eventID string, ticketID string, name string,
	price int32, desc string, stock int32, startTime time.Time, endTime time.Time) *UserTicketRepository.UserTicket {

	return &UserTicketRepository.UserTicket{
		EventID:    eventID,
		TicketID:   ticketID,
		Name:       name,
		Price:      price,
		Desc:       desc,
		Stock:      stock,
		StartTime:  startTime,
		EndTime:    endTime,
		CreateTime: txTime,
		UpdateTime: txTime,
	}
}

// CreateTicketInfo .
func CreateTicketInfo(userTicket *UserTicketRepository.UserTicket) *ticket.TicketInfo {

	if userTicket == nil {
		return nil
	}

	return &ticket.TicketInfo{
		TicketId:  userTicket.TicketID,
		EventId:   userTicket.EventID,
		Name:      userTicket.Name,
		Desc:      userTicket.Desc,
		Price:     userTicket.Price,
		Stock:     userTicket.Stock,
		StartTime: dateutil.TimeToTimestamp(userTicket.StartTime),
		EndTime:   dateutil.TimeToTimestamp(userTicket.EndTime),
	}
}

// CreateTicketPayTypeList .
func CreateTicketPayTypeList(mstTicketPayTypeList []*MstTicketPayTypeRepository.MstTicketPayType) []*ticket.TicketPayType {

	var ticketPayTypeList []*ticket.TicketPayType

	for _, mstTicketPayType := range mstTicketPayTypeList {

		ticketpayType := &ticket.TicketPayType{
			Id:   mstTicketPayType.ID,
			Name: mstTicketPayType.Name,
		}

		ticketPayTypeList = append(ticketPayTypeList, ticketpayType)
	}

	return ticketPayTypeList
}
