package UserTicketGroupRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_TICKET_GROUP"

	// FEventID .
	FEventID = "eid"

	fGroupName  = "name"
	fDesc       = "desc"
	fUpdateTime = "ut"
)

var reflectType = reflect.TypeOf((*UserTicketGroup)(nil))

// UserTicketGroup .
type UserTicketGroup struct {
	ID         *primitive.ObjectID `bson:"_id,omitempty"`
	GroupID    string              `bson:"gid"`
	GroupName  string              `bson:"name"`
	EventID    string              `bson:"eid"`
	Desc       string              `bson:"desc"`
	CreateTime time.Time           `bson:"ct"`
	UpdateTime time.Time           `bson:"ut"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert .
func Insert(ctx context.Context, userTicketGroup *UserTicketGroup) interface{} {
	return getDb(ctx).InsertOne(userTicketGroup)
}

// Update .
func Update(ctx context.Context, txTime time.Time, eventID string, groupID string, groupName string, desc string) int32 {

	query := bson.M{FEventID: eventID, groupID: groupID}

	update := bson.M{
		"$set": bson.M{
			"$set": bson.M{
				fUpdateTime: txTime,
				fGroupName:  groupName,
				fDesc:       desc,
			},
		},
	}

	return getDb(ctx).UpdateOne(query, update)
}

// Delete .
func Delete(ctx context.Context, eventID string, groupID string) int32 {
	query := bson.M{FEventID: eventID, groupID: groupID}

	return getDb(ctx).DeleteOne(query)
}

// Get .
func Get(ctx context.Context, eventID string, groupID string) *UserTicketGroup {
	query := bson.M{FEventID: eventID, groupID: groupID}

	result := getDb(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserTicketGroup)
}

// GetList .
func GetList(ctx context.Context, eventID string) []*UserTicketGroup {

	query := bson.M{FEventID: eventID}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserTicketGroup)
}
