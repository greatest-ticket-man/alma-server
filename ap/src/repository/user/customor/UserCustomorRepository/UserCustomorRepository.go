package UserCustomorRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// 顧客情報

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_CUSTOMOR"

	// FEmail .
	FEmail = "email"

	// FEventIDList .
	FEventIDList = "events"
)

var reflectType = reflect.TypeOf((*UserCustomor)(nil))

// UserCustomor .
// どの人たちまでが、この情報に手をつけられるかを決める, 組織ごとかな
type UserCustomor struct {
	ID                 string      `bson:"_id"`    // customorID
	Email              string      `bson:"email"`  //Emailと管理組織でuniqueにしようか考え中
	NameList           []*NameInfo `bson:"names"`  // 複数名前を登録する可能性があるため
	EventIDList        []string    `bson:"events"` // 参照できる、EventIDたち
	OrganizationIDList []string    `bson:"orgs"`   // 参照できる組織たち
	CreateTime         time.Time   `bson:"ct"`
	UpdateTime         time.Time   `bson:"ut"`
}

// NameInfo 複数の名前で登録した場合、判別が難しくなるため
// 今まで使用してきた名前を記録するようにする
type NameInfo struct {
	FirstName         string `bson:"fname"`
	FirstNameFurigana string `bson:"ff"`
	LastName          string `bson:"lname"`
	LastNameFurigana  string `bson:"lf"`
}

// getDb
func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// FindInEmail .
func FindInEmail(ctx context.Context, eventIDList []string, emailList []string) []*UserCustomor {

	query := bson.M{
		FEmail: bson.M{
			"$in": emailList,
		},
		FEventIDList: bson.M{
			"$in": eventIDList,
		},
	}

	result := getDb(ctx).Find(query, reflectType)
	if result == nil {
		return nil
	}

	return result.([]*UserCustomor)
}
