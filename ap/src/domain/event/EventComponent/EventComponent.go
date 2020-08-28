package EventComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"alma-server/ap/src/repository/user/event/UserEventInviteMemberRepository"
	"sort"
	"time"
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
			EventAuthDesc: mstEventAuth.Desc,
		}

		eventAuthInfoList = append(eventAuthInfoList, eventAuthInfo)
	}

	return eventAuthInfoList
}

// CreateInviteMemberList .
func CreateInviteMemberList(eventID string, txTime time.Time, list []*event.InviteMemberInfo) []*UserEventInviteMemberRepository.UserEventInviteMember {

	var userEventInviteMemberList []*UserEventInviteMemberRepository.UserEventInviteMember

	for _, inviteMemberInfo := range list {

		userEventInviteMember := &UserEventInviteMemberRepository.UserEventInviteMember{
			Email:      inviteMemberInfo.Email,
			EventID:    eventID,
			AuthID:     inviteMemberInfo.Authority,
			CreateTime: txTime,
			UpdateTime: txTime,
		}

		userEventInviteMemberList = append(userEventInviteMemberList, userEventInviteMember)

	}

	return userEventInviteMemberList
}
