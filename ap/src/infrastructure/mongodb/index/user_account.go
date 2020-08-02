package index

import (
	"alma-server/ap/src/repository/user/UserAccountRepository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	userIndexMap[UserAccountRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			Keys:    bson.M{UserAccountRepository.FEmail: 1},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
	}
}
