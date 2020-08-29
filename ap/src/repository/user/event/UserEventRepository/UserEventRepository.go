package UserEventRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// イベントの情報リポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_EVENT"

	// FEventID .
	FEventID = "_id"

	fName           = "name"
	fOrganization   = "organization"
	FMemberList     = "members"
	FTempMemberList = "tmps"
	fCreateTime     = "ct"
	fUpdateTime     = "ut"

	FMemberInfoMid = "mid"

	FTempMemberInfoEmail = "email"
)

var reflectType = reflect.TypeOf((*UserEvent)(nil))

// UserEvent .
type UserEvent struct {
	ID           string    `bson:"_id,omitempty"`
	Name         string    `bson:"name"`
	Organization string    `bson:"organization"`
	CreateTime   time.Time `bson:"ct"`
	UpdateTime   time.Time `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert イベントの作成
func Insert(ctx context.Context, txTime time.Time, eventID string, name string, organization string) interface{} {

	userEvent := &UserEvent{
		ID:           eventID,
		Name:         name,
		Organization: organization,
		CreateTime:   txTime,
		UpdateTime:   txTime,
	}

	return getDb(ctx).InsertOne(userEvent)
}

// Remove イベントの削除
func Remove(ctx context.Context, eventID string) int32 {
	query := bson.M{FEventID: eventID}
	return getDb(ctx).DeleteOne(query)
}

// Update イベントの編集
func Update(ctx context.Context, txTime time.Time, eventID string, name string, organization string) int32 {

	query := bson.M{FEventID: eventID}

	update := bson.M{
		"$set": bson.M{
			fName:         name,
			fOrganization: organization,
			fUpdateTime:   txTime,
		},
	}

	return getDb(ctx).UpdateOne(query, update)
}

// FindOneAndUpdate イベントを取得してその後Updateする
func FindOneAndUpdate(ctx context.Context, txTime time.Time, eventID string, name string, organization string) *UserEvent {

	query := bson.M{FEventID: eventID}

	update := bson.M{
		"$set": bson.M{
			fName:         name,
			fOrganization: organization,
			fUpdateTime:   txTime,
		},
	}

	result := getDb(ctx).FindOneAndUpdate(query, update, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserEvent)
}

// Get .
func Get(ctx context.Context, eventID string) *UserEvent {

	query := bson.M{FEventID: eventID}

	result := getDb(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserEvent)
}

// GetListInEventID .
func GetListInEventID(ctx context.Context, eventIDList []string) []*UserEvent {

	query := bson.M{
		FEventID: bson.M{
			"$in": eventIDList,
		},
	}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserEvent)
}
