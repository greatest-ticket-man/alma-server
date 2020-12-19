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
	ID         string    `bson:"_id"`
	Seq        uint64    `bson:"seq"` // Eventでの範囲の番号
	EventID    string    `bson:"eid"`
	CustomorID string    `bson:"cid"` // お客さん情報
	Name       string    `bson:"name"`
	Furigana   string    `bson:"furigana"`
	Email      string    `bson:"email"`
	TicketID   string    `bson:"tid"`  // 購入したチケットID
	TicketCnt  int32     `bson:"tcnt"` // 何枚購入したか
	PayTypeID  string    `bson:"ptid"` // MstTicketPayType
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert .
func Insert(ctx context.Context, txTime time.Time, reserveID string, seq uint64, eventID string, customorID string, name string, furigana string, email string, ticketID string, ticketCnt int32, scheduleID string, payTypeID string) interface{} {

	userReserve := &UserReserve{
		ID:         reserveID,
		Seq:        seq,
		EventID:    eventID,
		CustomorID: customorID,
		Name:       name,
		Furigana:   furigana,
		Email:      email,
		TicketID:   ticketID,
		TicketCnt:  ticketCnt,
		CreateTime: txTime,
		UpdateTime: txTime,
		PayTypeID:  payTypeID,
	}

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
