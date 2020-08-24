package UserEventRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"gopkg.in/mgo.v2/bson"
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
	ID             string            `bson:"_id,omitempty"`
	Name           string            `bson:"name"`
	Organization   string            `bson:"organization"`
	MemberList     []*MemberInfo     `bson:"members"`
	TempMemberList []*TempMemberInfo `bson:"tmps"`
	CreateTime     time.Time         `bson:"ct"`
	UpdateTime     time.Time         `bson:"ut"`
}

// MemberInfo .
type MemberInfo struct {
	Mid        string    `bson:"mid"`
	AuthID     string    `bson:"auth"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

// TempMemberInfo .
type TempMemberInfo struct {
	Email      string    `bson:"email"`
	AuthID     string    `bson:"auth"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert イベントの作成
func Insert(ctx context.Context, txTime time.Time, eventID string, name string, organization string, memberList []*MemberInfo, tempMemberList []*TempMemberInfo) interface{} {

	userEvent := &UserEvent{
		ID:             eventID,
		Name:           name,
		Organization:   organization,
		MemberList:     memberList,
		TempMemberList: tempMemberList,
		CreateTime:     txTime,
		UpdateTime:     txTime,
	}

	return getDb(ctx).InsertOne(userEvent)
}

// Update イベントの編集
func Update(ctx context.Context, txTime time.Time, eventID string, name string, organization string, tempMemberList []*TempMemberInfo) int32 {

	query := bson.M{FEventID: eventID}

	update := bson.M{
		"$set": bson.M{
			fName:         name,
			fOrganization: organization,
			fUpdateTime:   txTime,
		},
		"$addToSet": bson.M{
			FTempMemberList: bson.M{
				"$each": tempMemberList,
			},
		},
	}

	return getDb(ctx).UpdateOne(query, update)
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
