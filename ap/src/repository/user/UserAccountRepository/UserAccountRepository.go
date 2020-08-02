package UserAccountRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ユーザーのアカウント情報のリポジトリ
const (
	// ThisCollectionName .
	ThisCollectionName = "USER_ACCOUNT"
)

var reflectType = reflect.TypeOf(&UserAccount{})

// UserAccount .
type UserAccount struct {

	ID string `bson:"_id,omitempty"` // mid
	Email string `bson:"email"`
	Password string `bson:"pass"`
	Use bool `bson:"use"`

	LoginTime time.Time `bson:"lt"`
	CreateTime time.Time `bson:"ct"`
	UpdateTime time.Time `bson:"ut"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Insert .
func Insert(ctx context.Context, userAccount *UserAccount) string {

	

}
