package mongodb

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"context"
	"fmt"
	"net/url"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonoptions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/mongo/driver/connstring"
)

var clientMap map[string]*mongo.Client

// UserDbName .
var UserDbName = "ALMA_USER"

// LogDbName .
var LogDbName = "ALMA_LOG"

// Setup MongoDB Connect
func Setup(configList []*config.MongoDB) bool {

	clientMap = map[string]*mongo.Client{}
	for _, config := range configList {

		uri := createApplyURI(config)
		client := connect(uri)

		clientMap[config.Db] = client
	}

	return true
}

// createApplyURI MongoDBのURIを作成する
func createApplyURI(config *config.MongoDB) string {

	option := config.Option

	value := url.Values{}
	value.Add("connecttimeoutms", option.ConnectTimeoutMs)
	value.Add("heartbeatintervalms", option.HeartBeatIntervalMs)
	value.Add("maxidletimems", option.MaxIdleTimeMs)
	value.Add("maxpoolsize", option.MaxPoolSize)
	value.Add("readpreference", option.ReadPreference)
	value.Add("readconcernlevel", option.ReadConcernLevel)
	value.Add("w", option.WriteConnection)
	value.Add("sockettimeoutms", option.SocketTimeoutMs)

	// userが空白でなければ、追加する
	db := ""
	var user *url.Userinfo
	if config.User != "" {
		user = url.UserPassword(config.User, config.Password)
		db = config.Db
	}

	// make url
	u := url.URL{
		Scheme:   "mongodb",
		User:     user,
		Host:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Path:     fmt.Sprintf("/%s", db),
		RawQuery: value.Encode(),
	}

	// validateion check
	uri := u.String()
	connString, err := connstring.Parse(uri)
	chk.SE(err)

	if len(connString.UnknownOptions) != 0 {
		chk.SE(err, "MongoDB接続OptionでUnknownOptionsがみつかりました:"+uri)
	}

	return uri
}

// connect clientを生成する
func connect(uri string) *mongo.Client {

	opt := options.Client().ApplyURI(uri)
	chk.SE(opt.Validate())

	// time.TimeをLocalTimeZoneで取得できるようにする
	rb := bson.NewRegistryBuilder()
	rb.RegisterTypeDecoder(
		reflect.TypeOf(time.Time{}),
		bsoncodec.NewTimeCodec(
			bsonoptions.TimeCodec().SetUseLocalTimeZone(true),
		),
	)
	opt.SetRegistry(rb.Build())

	client, err := mongo.Connect(context.Background(), opt)
	chk.SE(err, "MongoDBのConnectに失敗:"+uri)

	// pingで接続許可されていることを確認する
	chk.SE(client.Ping(context.Background(), readpref.Primary()), "MongoDBにPingできませんでした:"+uri)

	return client
}

// GetUserCollection get user collection
func GetUserCollection(ctx context.Context, colName string) *AlmaCollection {
	return GetDbCollection(ctx, UserDbName, colName)
}

// GetLogCollection get log collection
func GetLogCollection(ctx context.Context, colName string) *AlmaCollection {
	return GetDbCollection(ctx, LogDbName, colName)
}

// GetDbCollection get any collection
func GetDbCollection(ctx context.Context, dbName string, colName string) *AlmaCollection {
	return &AlmaCollection{
		col: clientMap[dbName].Database(dbName).Collection(colName),
		ctx: ctx,
	}
}
