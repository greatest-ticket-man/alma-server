package UserAccountRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// ユーザーのアカウント情報のリポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_ACCOUNT"

	FEmail = "email"
	FMid   = "_id"

	fLoginTime = "lt"
)

var reflectType = reflect.TypeOf(&UserAccount{})

// UserAccount .
type UserAccount struct {
	ID       string `bson:"_id,omitempty"` // mid
	Name     string `bson:"name"`          // 名前
	Email    string `bson:"email"`
	Password string `bson:"pass"` // md5 hashされたpassword
	Use      bool   `bson:"use"`

	LoginTime  time.Time `bson:"lt"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert アカウントの作成
func Insert(ctx context.Context, txTime time.Time, mid string, name string, email string, password string) string {

	userAccount := &UserAccount{
		ID:         mid,
		Name:       name,
		Email:      email,
		Password:   password,
		Use:        true,
		LoginTime:  txTime,
		CreateTime: txTime,
		UpdateTime: txTime,
	}

	result := getDb(ctx).InsertOne(userAccount)
	return result.(string)
}

// GetFromEmail メールアドレスからデータを取得する
func GetFromEmail(ctx context.Context, email string) *UserAccount {

	query := bson.M{FEmail: email}

	result := getDb(ctx).FindOne(query, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserAccount)
}

// GetList midListからデータを取得する
func GetList(ctx context.Context, midList []string) []*UserAccount {

	query := bson.M{
		FMid: bson.M{
			"$in": midList,
		},
	}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserAccount)
}

// FindAndUpdate .
func FindAndUpdate(ctx context.Context, email string, txTime time.Time) *UserAccount {

	query := bson.M{FEmail: email}

	update := bson.M{
		"$set": bson.M{
			fLoginTime: txTime,
		},
	}

	result := getDb(ctx).FindOneAndUpdate(query, update, reflectType)
	if result == nil {
		return nil
	}

	return result.(*UserAccount)
}
