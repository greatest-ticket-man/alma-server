package UserTicketRepositoryV2

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_TICKET_V2"
)

var reflectType = reflect.TypeOf((*UserTicketV2)(nil))

// UserTicketV2 .
type UserTicketV2 struct {
	ID             *primitive.ObjectID         `bson:"_id,omitempty"`
	TicketID       string                      `bson:"tid"`
	EventID        string                      `bson:"eid"`
	Name           string                      `bson:"name"`
	Desc           string                      `bson:"desc"`
	BuyStartTime   time.Time                   `bson:"bst"` // チケット
	BuyEndTime     time.Time                   `bson:"bet"` // チケット販売終了時間
	EventStartTime time.Time                   `bson:"est"`
	TicketGroupID  *primitive.ObjectID         `bson:"tgid"`
	PricesMap      map[string]*TicketPriceInfo `bson:"prices"`
	CreateTime     time.Time                   `bson:"ct"`
	UpdateTime     time.Time                   `bson:"ut"`
}

// TicketPriceInfo 価格帯設定
type TicketPriceInfo struct {
	PriceID    string    `bson:"id"`
	Name       string    `bson:"name"`
	PriceType  int32     `bson:"pt"` // 価格のタイプ: 前払い/当日生産/クレジット
	Price      int32     `bson:"price"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

// getDb .
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}
