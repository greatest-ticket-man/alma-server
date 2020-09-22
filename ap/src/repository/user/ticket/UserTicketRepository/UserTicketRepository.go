package UserTicketRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ユーザーが作成したチケット情報のリポジトリ

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_TICKET"

	// FTicketID .
	FTicketID = "tid"

	// FEventID .
	FEventID = "eid"

	fName       = "name"
	fDesc       = "desc"
	fPrice      = "price"
	fCreateTime = "ct"
	fUpdateTime = "ut"
)

var reflectType = reflect.TypeOf((*UserTicket)(nil))

// UserTicket TicketIDとEventIDでUniqueになる予定
type UserTicket struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	TicketID   string              `bson:"tid"`
	EventID    string              `bson:"eid"` // イベントのID
	Name       string              `bson:"name"`
	Desc       string              `bson:"desc"` // チケットの説明
	Price      int32               `bson:"price"`
	CreateTime time.Time           `bson:"ct"`
	UpdateTime time.Time           `bson:"ut"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// FindOne チケット情報を一つ取得する
func FindOne(ctx context.Context, eventID string, ticketID string) *UserTicket {

	query := bson.M{FEventID: eventID, FTicketID: ticketID}

	result := getDb(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserTicket)
}

// Find イベントに紐づいたチケットを全て取得する
func Find(ctx context.Context, eventID string) []*UserTicket {

	query := bson.M{FEventID: eventID}
	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserTicket)
}

// Insert チケットの追加
func Insert(ctx context.Context, txTime time.Time, ticketID string, eventID string,
	name string, desc string, price int32) interface{} {

	userTicket := &UserTicket{
		TicketID:   ticketID,
		EventID:    eventID,
		Name:       name,
		Desc:       desc,
		Price:      price,
		CreateTime: txTime,
		UpdateTime: txTime,
	}

	return getDb(ctx).InsertOne(userTicket)
}

// Update チケット情報の更新
func Update(ctx context.Context, txTime time.Time, beforeTicketID string, eventID string,
	name string, desc string, price int32, ticketID string) int32 {

	query := bson.M{FEventID: eventID, FTicketID: beforeTicketID}

	update := bson.M{
		"$set": bson.M{
			fName:       name,
			FTicketID:   ticketID,
			fDesc:       desc,
			fPrice:      price,
			fUpdateTime: txTime,
		},
	}

	return getDb(ctx).UpdateOne(query, update)
}

// Remove チケットの削除
func Remove(ctx context.Context, ticketID string, eventID string) int32 {
	query := bson.M{FEventID: eventID, FTicketID: ticketID}
	return getDb(ctx).DeleteOne(query)
}
