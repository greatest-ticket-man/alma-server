package UserEventRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"
)

// イベントの情報リポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT"
)

var reflectType = reflect.TypeOf(&UserEvent{})

// UserEvent .
type UserEvent struct {
	ID            string            `bson:"_id,omitempty"`
	Name          string            `bson:"name"`
	Organization  string            `bson:"organization"`
	MemberMap     map[string]string `bson:"members"`
	TempMemberMap map[string]string `bson:"tempmembers"`
	CreateTime    time.Time         `bson:"ct"`
	UpdateTime    time.Time         `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert イベントの作成
func Insert(ctx context.Context, txTime time.Time, eventID string, name string, organization string, memberMap map[string]string, tempMemberMap map[string]string) interface{} {

	userEvent := &UserEvent{
		ID:           eventID,
		Name:         name,
		Organization: organization,
		CreateTime:   txTime,
		UpdateTime:   txTime,
	}

	return getDb(ctx).InsertOne(userEvent)
}
