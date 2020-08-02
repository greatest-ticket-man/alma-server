package UserAccountRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"
)

// ユーザーのアカウント情報のリポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_ACCOUNT"

	FEmail = "email"
)

var reflectType = reflect.TypeOf(&UserAccount{})

// UserAccount .
type UserAccount struct {
	ID       string `bson:"_id,omitempty"` // mid
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
func Insert(ctx context.Context, txTime time.Time, mid string, email string, password string) string {

	userAccount := &UserAccount{
		ID:         mid,
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
