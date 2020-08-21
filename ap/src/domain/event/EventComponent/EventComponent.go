package EventComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"sort"
)

// CreateEventAuthInfoList .
func CreateEventAuthInfoList(mstEventAuthList []*MstEventAuthRepository.MstEventAuth) []*event.EventAuthInfo {

	// 先に表示順でソート
	sort.Slice(mstEventAuthList, func(i, j int) bool { return mstEventAuthList[i].Order < mstEventAuthList[j].Order })

	var eventAuthInfoList []*event.EventAuthInfo

	for _, mstEventAuth := range mstEventAuthList {

		eventAuthInfo := &event.EventAuthInfo{
			EventAuthId:   mstEventAuth.ID,
			EventAuthName: mstEventAuth.Name,
			EventAtuhDesc: mstEventAuth.Desc,
		}

		eventAuthInfoList = append(eventAuthInfoList, eventAuthInfo)
	}

	return eventAuthInfoList
}
