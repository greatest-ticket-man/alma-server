package ReserveComponent

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
	"alma-server/ap/src/repository/user/reserve/UserReserveRepository"
)

// CreateReserveInfoList .
func CreateReserveInfoList(userReserveList []*UserReserveRepository.UserReserve) []*reserve.ReserveInfo {

	var reserveInfoList []*reserve.ReserveInfo

	for _, userReserve := range userReserveList {

		reserveInfo := &reserve.ReserveInfo{
			ReserveId:    userReserve.ID,
			No:           0, // TODO
			TicketName:   userReserve.TicketID + ":TODO",
			TicketNum:    userReserve.TicketCnt,
			Name:         userReserve.Name,
			NameFurigana: userReserve.Furigana,
			Email:        "TODO: customorから取得する",
			CustomorId:   "TODO: customorができてから追加する",
			CreatedAt:    dateutil.TimeToTimestamp(userReserve.CreateTime),
			UpdatedAt:    dateutil.TimeToTimestamp(userReserve.UpdateTime),
		}

		reserveInfoList = append(reserveInfoList, reserveInfo)
	}

	return reserveInfoList
}
