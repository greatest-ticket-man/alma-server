package UserReserveRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 予約情報を記録するリポジトリ

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_RESERVE"

	// FReserveID .
	FReserveID = "_id"

	// FEventID .
	FEventID = "eid"
)

var reflectType = reflect.TypeOf((*UserReserve)(nil))

// UserReserve .
type UserReserve struct {
	ID                string    `bson:"_id"`
	No                int32     `bson:"no"` // Eventでの範囲の番号
	EventID           string    `bson:"eid"`
	CustomorID        string    `bson:"cid"` // お客さん情報
	FirstName         string    `bson:"fname"`
	FirstNameFurigana string    `bson:"ff"`
	LastName          string    `bson:"lname"`
	LastNameFurigana  string    `bson:"lf"`
	TicketID          string    `bson:"tid"`  // 購入したチケットID
	TicketCnt         int32     `bson:"tcnt"` // 何枚購入したか
	CreateTime        time.Time `bson:"ct"`
	UpdateTime        time.Time `bson:"ut"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Add .
func Add(ctx context.Context, userReserve *UserReserve) interface{} {
	return getDb(ctx).InsertOne(userReserve)
}

// GetList 条件が増えるかもなので、Findで
func GetList(ctx context.Context, eventID string) []*UserReserve {

	query := bson.M{FEventID: eventID}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserReserve)
}

// Get 予約情報を一つ取得する
func Get(ctx context.Context, eventID string, reserveID string) *UserReserve {

	query := bson.M{FReserveID: reserveID, FEventID: eventID}

	result := getDb(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserReserve)
}
