package UserEventRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// イベントの情報リポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT"
)

var reflectType = reflect.TypeOf(&UserEvent{})

// UserEvent .
type UserEvent struct {
	ID            *primitive.ObjectID `bson:"_id,omitempty"` // ObjectID
	Name          string              `bson:"name"`
	Organization  string              `bson:"organization"`
	MemberMap     map[string]string   `bson:"members"`
	TempMemberMap map[string]string   `bson:"tempmembers"`
	CreateTime    time.Time           `bson:"ct"`
	UpdateTime    time.Time           `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert イベントの作成
func Insert(ctx context.Context, txTime time.Time, name string, organization string, memberMap map[string]string, tempMemberMap map[string]string) *primitive.ObjectID {

	userEvent := &UserEvent{
		Name:         name,
		Organization: organization,
		CreateTime:   txTime,
		UpdateTime:   txTime,
	}

	result := getDb(ctx).InsertOne(userEvent).(primitive.ObjectID)
	return &result
}
