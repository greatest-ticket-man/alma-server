package UserEventInviteMemberRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// イベントの招待メンバーリポジトリ

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT_INVITE_MEMBER"

	// FEmail .
	FEmail = "_id"
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

// RemoveMany 複数を削除
func RemoveMany(ctx context.Context, emailList []string) int32 {

	query := bson.M{FEmail: bson.M{
		"$in": emailList,
	}}

	return getDb(ctx).DeleteMany(query)
}

// toInterface sliceをinterfaceに変換する
func toInterface(userEventInviteMemberList []*UserEventInviteMember) []interface{} {
	var list []interface{}
	for _, userEventInviteMember := range userEventInviteMemberList {
		list = append(list, userEventInviteMember)
	}
	return list
}
