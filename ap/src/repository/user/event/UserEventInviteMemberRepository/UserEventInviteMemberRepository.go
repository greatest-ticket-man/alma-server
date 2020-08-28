package UserEventInviteMemberRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"
)

// イベントの招待メンバーリポジトリ

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT_INVITE_MEMBER"
)

var reflectType = reflect.TypeOf((*UserEventInviteMember)(nil))

// UserEventInviteMember イベントのメンバーへの招待中のメンバー
type UserEventInviteMember struct {
	Email      string    `bson:"_id"` // email
	EventID    string    `bson:"event"`
	AuthID     string    `bson:"auth"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// InsertBulk 一括で招待メンバーを登録する
func InsertBulk(ctx context.Context, userEventInviteMemberList []*UserEventInviteMember) []interface{} {
	return getDb(ctx).InsertMany(toInterface(userEventInviteMemberList))
}

// toInterface sliceをinterfaceに変換する
func toInterface(userEventInviteMemberList []*UserEventInviteMember) []interface{} {
	var list []interface{}
	for _, userEventInviteMember := range userEventInviteMemberList {
		list = append(list, userEventInviteMember)
	}
	return list
}
