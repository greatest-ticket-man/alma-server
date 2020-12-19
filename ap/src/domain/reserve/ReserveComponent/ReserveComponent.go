package ReserveComponent

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
	"alma-server/ap/src/repository/master/ticket/MstTicketPayTypeRepository"
	"alma-server/ap/src/repository/user/reserve/UserReserveRepository"
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"
)

// CreateReserveInfoList .
func CreateReserveInfoList(userReserveList []*UserReserveRepository.UserReserve,
	userTicketMap map[string]*UserTicketRepository.UserTicket,
	mstTicketPayTypeMap map[string]*MstTicketPayTypeRepository.MstTicketPayType) []*reserve.ReserveInfo {

	var reserveInfoList []*reserve.ReserveInfo

	for _, userReserve := range userReserveList {

		userTicket := userTicketMap[userReserve.TicketID]

		reserveInfo := &reserve.ReserveInfo{
			ReserveId:      userReserve.ID,
			Seq:            userReserve.Seq,
			TicketName:     userTicket.Name,
			TicketNum:      userReserve.TicketCnt,
			EventStartTime: dateutil.TimeToTimestamp(userTicket.StartTime),
			Name:           userReserve.Name,
			NameFurigana:   userReserve.Furigana,
			Email:          userReserve.Email,
			PayTypeName:    mstTicketPayTypeMap[userReserve.PayTypeID].Name,
			CustomorId:     userReserve.CustomorID,
			CreatedAt:      dateutil.TimeToTimestamp(userReserve.CreateTime),
			UpdatedAt:      dateutil.TimeToTimestamp(userReserve.UpdateTime),
		}

		reserveInfoList = append(reserveInfoList, reserveInfo)
	}

	return reserveInfoList
}
