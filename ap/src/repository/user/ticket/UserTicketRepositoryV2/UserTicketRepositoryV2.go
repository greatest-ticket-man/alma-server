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
	ID             *primitive.ObjectID `bson:"_id,omitempty"`
	TicketID       string              `bson:"tid"`
	EventID        string              `bson:"eid"`
	Name           string              `bson:"name"`
	Desc           string              `bson:"desc"`
	EventStartTime time.Time           `bson:"est"`
	TicketGroupID  *primitive.ObjectID `bson:"tgid"`
	CreateTime     time.Time           `bson:"ct"`
	UpdateTime     time.Time           `bson:"ut"`
}

// getDb .
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}
