package UserEventRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"fmt"
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

	fName          = "name"
	fOrganization  = "organization"
	fMemberMap     = "members"
	fTempMemberMap = "tempmembers"
	fCreateTime    = "ct"
	fUpdateTime    = "ut"
)

var reflectType = reflect.TypeOf((*UserEvent)(nil))

// var reflectType = reflect.TypeOf(&UserEvent{})

// UserEvent .
type UserEvent struct {
	ID            string            `bson:"_id,omitempty"`
	Name          string            `bson:"name"`
	Organization  string            `bson:"organization"`
	MemberMap     map[string]string `bson:"members,omitempty"` // https://docs.mongodb.com/manual/core/index-multikey/ マルチキーインデックスを使って、主t句できるようにする // 将来的にListにする？
	TempMemberMap map[string]string `bson:"tempmembers"`       // key: email, value: 権限
	CreateTime    time.Time         `bson:"ct"`
	UpdateTime    time.Time         `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert イベントの作成
func Insert(ctx context.Context, txTime time.Time, eventID string, name string, organization string, memberMap map[string]string, tempMemberMap map[string]string) interface{} {

	userEvent := &UserEvent{
		ID:            eventID,
		Name:          name,
		Organization:  organization,
		MemberMap:     memberMap,
		TempMemberMap: tempMemberMap,
		CreateTime:    txTime,
		UpdateTime:    txTime,
	}

	return getDb(ctx).InsertOne(userEvent)
}

// Update イベントの編集
func Update(ctx context.Context, txTime time.Time, eventID string, name string, organization string, memberMap map[string]string, tempMemberMap map[string]string) int32 {

	query := bson.M{FEventID: eventID}

	set := bson.M{
		fName:         name,
		fOrganization: organization,
		fUpdateTime:   txTime,
	}

	// tempMemberMap
	for email, auth := range tempMemberMap {
		set[fmt.Sprintf("%s.%s", fTempMemberMap, email)] = auth // TODO .をescapeする方法を探す
	}

	update := bson.M{
		"$set": set,
	}

	// TODO memberMap
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
