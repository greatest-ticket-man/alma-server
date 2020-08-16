package EventRpcService

import (
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"context"
	"time"
)

// CreateEvent .
func CreateEvent(ctx context.Context, mid string, txTime time.Time, eventName string, organizationName string, memberInfoList []*event.MemberInfo) *event.CreateEventReply {

	// log.Println("mid", mid, "txTime", txTime, "eventName", eventName, "organizationName", organizationName, "memberInfoList is ", jsonutil.Marshal(memberInfoList))

	// organizationは存在しない場合はLogicError

	// memberは、メールだけ送る。ペンディング状態にする。メールで参加したら本登録になる

	eventID := uniqueidutil.GenerateUniqueID()

	// 追加
	UserEventRepository.Insert(ctx, txTime, eventID, eventName, organizationName, map[string]string{"": ""}, map[string]string{"": ""})

	return &event.CreateEventReply{
		EventId:   eventID,
		EventName: eventName,
	}
}
