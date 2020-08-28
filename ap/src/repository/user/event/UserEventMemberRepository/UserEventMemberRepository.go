package UserEventMemberRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// イベントのメンバーリポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT_MEMBER"

	fMemberID = "_id"

	FEventID = "event"
	FMid     = "mid"
	fAuth    = "auth"

	fCreateTime = "ct"
	fUpdateTime = "ut"
)

var reflectType = reflect.TypeOf((*UserEventMember)(nil))

// UserEventMember .
type UserEventMember struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	EventID    string              `bson:"event"`
	Mid        string              `bson:"mid"`
	AuthID     string              `bson:"auth"`
	CreateTime time.Time           `bson:"ct"`
	UpdateTime time.Time           `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// InsertBulk メンバーを複数追加する
func InsertBulk(ctx context.Context, txTime time.Time, userEventMemberList []*UserEventMember) []interface{} {
	return getDb(ctx).InsertMany(toInterface(userEventMemberList))
}

// Upsert メンバーを一人追加する、すでにいる場合はAuthIDのみUpdateする
func Upsert(ctx context.Context, txTime time.Time, mid string, eventID string, authID string) int32 {

	query := bson.M{FEventID: eventID, FMid: mid}

	upsert := bson.M{
		"$setOnInsert": bson.M{
			FMid:        mid,
			FEventID:    eventID,
			fCreateTime: txTime,
		},
		"$set": bson.M{
			fAuth:       authID,
			fUpdateTime: txTime,
		},
	}

	return getDb(ctx).UpsertOne(query, upsert)
}

// RemoveMany メンバーを複数削除する
func RemoveMany(ctx context.Context, memberIDList []*primitive.ObjectID) int32 {

	query := bson.M{
		fMemberID: bson.M{
			"$in": memberIDList,
		},
	}

	return getDb(ctx).DeleteMany(query)
}

// toInterface sliceをintertfaceに変換する
func toInterface(userEventMemberList []*UserEventMember) []interface{} {

	var list []interface{}
	for _, userEventMember := range userEventMemberList {
		list = append(list, userEventMember)
	}

	return list
}
