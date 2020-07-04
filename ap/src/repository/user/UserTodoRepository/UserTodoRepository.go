package UserTodoRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
)

// ThisCollectionName .
const ThisCollectionName = "USER_TODO"

var reflectType = reflect.TypeOf(&USER_TODO{})

// USER_TODO やることをリスト
type USER_TODO struct {
	ID           *primitive.ObjectID `bson:"_id,omitempty"`
	Name         string              `bson:"name"`
	Title        string              `bson:"title"`
	Desc         string              `bson:"desc"`
	DeadlineTime time.Time           `bson:"dt"`
	CreateTime   time.Time           `bson:"ct"`
	UpdateTime   time.Time           `bson:"ut"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert .
func Insert(ctx context.Context, userTodo *USER_TODO) bool {

	getDb(ctx).InsertOne(userTodo)
	return true
}

// Delete .
func Delete(ctx context.Context, id *primitive.ObjectID) int32 {
	return getDb(ctx).DeleteOne(bson.M{"_id": id})
}

// Find .
func Find(ctx context.Context, name string) []*USER_TODO {
	query := bson.M{"name": name}
	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*USER_TODO)
}
