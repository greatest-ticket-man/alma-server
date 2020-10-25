package UserTicketRepositoryV2

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_TICKET_V2"
)

var reflectType = reflect.TypeOf((*UserTicketV2)(nil))

// UserTicketV2 .
type UserTicketV2 struct {
	ID *primitive.ObjectID `bson:""`
}

// getDb .
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}
