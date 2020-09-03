package MemberComponent

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/domain/account/AccountComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/member"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"alma-server/ap/src/repository/user/event/UserEventMemberRepository"
)

// GetMidListFromUserEventMember .
func GetMidListFromUserEventMember(userEventMemberList []*UserEventMemberRepository.UserEventMember) []string {
	var midList []string
	for _, userEventMember := range userEventMemberList {
		midList = append(midList, userEventMember.Mid)
	}

	return midList
}

// CreateMemberInfoList .
func CreateMemberInfoList(userEventMemberList []*UserEventMemberRepository.UserEventMember, userAccountList []*UserAccountRepository.UserAccount) []*member.MemberInfo {

	userAccountMap := AccountComponent.GetUserAccountMap(userAccountList)

	var memberInfoList []*member.MemberInfo

	for _, userEventMember := range userEventMemberList {

		userAccount := userAccountMap[userEventMember.Mid]
		if userAccount == nil {
			continue
		}

		memberInfo := &member.MemberInfo{
			Name:      userAccount.Name,
			Email:     userAccount.Email,
			Auth:      userEventMember.AuthID, // TODO AuthNameにするかも
			CreatedAt: dateutil.TimeToTimestamp(userEventMember.CreateTime),
		}

		memberInfoList = append(memberInfoList, memberInfo)

	}

	return memberInfoList
}
