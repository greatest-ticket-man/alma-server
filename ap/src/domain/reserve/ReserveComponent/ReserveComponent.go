package ReserveComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
	"alma-server/ap/src/repository/user/reserve/UserReserveRepository"
)

// CreateReserveInfoList .
func CreateReserveInfoList(userReserveList []*UserReserveRepository.UserReserve) []*reserve.ReserveInfo {

	var reserveInfoList []*reserve.ReserveInfo

	for _, userReserve := range userReserveList {

		reserveInfo := &reserve.ReserveInfo{
			ReserveId:         userReserve.ID,
			No:                0, // TODO
			TicketName:        userReserve.TicketID + ":TODO",
			TicketNum:         userReserve.TicketCnt,
			FirstName:         userReserve.FirstName,
			FirstNameFurigana: userReserve.FirstNameFurigana,
			LastName:          userReserve.LastName,
			LastNameFurigana:  userReserve.LastNameFurigana,
			Notes:             "TODO",
			Email:             "TODO: customorから取得する",
			CustomorId:        "TODO: customorができてから追加する",
		}

		reserveInfoList = append(reserveInfoList, reserveInfo)
	}

	return reserveInfoList
}
