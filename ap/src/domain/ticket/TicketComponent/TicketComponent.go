package TicketComponent

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/infrastructure/grpc/proto/ticket"
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
func CreateUserTicket(txTime time.Time, eventID string, ticketID string, name string, price int32, desc string, scheduleStockInfoList []*ticket.TicketScheduleStockInfo) *UserTicketRepository.UserTicket {

	scheduleStockMap := map[string]*UserTicketRepository.ScheduleStockInfo{}

	for _, scueduleStockInfo := range scheduleStockInfoList {

		scheduleStockMap[uniqueidutil.GenerateUniqueID()] = &UserTicketRepository.ScheduleStockInfo{
			EventStartTime: dateutil.TimestampToTime(scueduleStockInfo.EventStartTime),
			Stock:          scueduleStockInfo.Stock,
		}

	}

	return &UserTicketRepository.UserTicket{
		EventID:              eventID,
		TicketID:             ticketID,
		Name:                 name,
		Price:                price,
		Desc:                 desc,
		ScheduleStockInfoMap: scheduleStockMap,
		CreateTime:           txTime,
		UpdateTime:           txTime,
	}
}

// CreateTicketInfo .
func CreateTicketInfo(userTicket *UserTicketRepository.UserTicket) *ticket.TicketInfo {

	if userTicket == nil {
		return nil
	}

	var scheduleStockList []*ticket.TicketScheduleStockInfo
	for id, scheduleStockInfo := range userTicket.ScheduleStockInfoMap {
		scheduleStockList = append(scheduleStockList, &ticket.TicketScheduleStockInfo{
			ScheduleStockId: id,
			EventStartTime:  dateutil.TimeToTimestamp(scheduleStockInfo.EventStartTime),
			Stock:           scheduleStockInfo.Stock,
		})
	}

	return &ticket.TicketInfo{
		TicketId:          userTicket.TicketID,
		EventId:           userTicket.EventID,
		Name:              userTicket.Name,
		Desc:              userTicket.Desc,
		Price:             userTicket.Price,
		ScheduleStockList: scheduleStockList,
	}
}

// CreasteScheduleStockMap .
func CreasteScheduleStockMap(scheduleStockInfoList []*ticket.TicketScheduleStockInfo) map[string]*UserTicketRepository.ScheduleStockInfo {

	scheduleStockMap := map[string]*UserTicketRepository.ScheduleStockInfo{}
	for _, scheduleStockInfo := range scheduleStockInfoList {

		scheduleStock := &UserTicketRepository.ScheduleStockInfo{
			EventStartTime: dateutil.TimestampToTime(scheduleStockInfo.EventStartTime),
			Stock:          scheduleStockInfo.Stock,
		}
		scheduleStockMap[scheduleStockInfo.ScheduleStockId] = scheduleStock
	}

	return scheduleStockMap
}
