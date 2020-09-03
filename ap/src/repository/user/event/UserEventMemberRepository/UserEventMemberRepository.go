package UserEventMemberRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// GetList 自分が参加しているイベントデータを取得する
func GetList(ctx context.Context, mid string) []*UserEventMember {

	query := bson.M{FMid: mid}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserEventMember)
}

// GetListFromEventID 指定したイベントに参加している人を取得する
func GetListFromEventID(ctx context.Context, eventID string) []*UserEventMember {

	query := bson.M{FEventID: eventID}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserEventMember)
}

// InsertBulk メンバーを複数追加する
func InsertBulk(ctx context.Context, txTime time.Time, userEventMemberList []*UserEventMember) []interface{} {
	return getDb(ctx).InsertMany(toInterface(userEventMemberList))
}

// Upsert メンバーを一人追加する、すでにいる場合はAuthIDのみUpdateする
func Upsert(ctx context.Context, mid string, txTime time.Time, eventID string, authID string) int32 {

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

// FindOneAndUpsert 取得して、メンバー情報を変更する
// 変更前のデータを取得します
func FindOneAndUpsert(ctx context.Context, mid string, txTime time.Time, eventID string, authID string) *UserEventMember {

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

	isBefore := options.Before
	isUpsert := true

	opt := &options.FindOneAndUpdateOptions{
		ReturnDocument: &isBefore,
		Upsert:         &isUpsert,
	}

	result := getDb(ctx).FindOneAndUpdate(query, upsert, reflectType, opt)
	if result == nil {
		return nil
	}

	return result.(*UserEventMember)
}

// Remove メンバーを一人削除する
func Remove(ctx context.Context, mid string, txTime time.Time, eventID string) int32 {
	query := bson.M{FMid: mid, FEventID: eventID}

	return getDb(ctx).DeleteOne(query)
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
