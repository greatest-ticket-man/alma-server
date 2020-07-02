package index

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	userIndexMap["USER_TICKET"] = []mongo.IndexModel{
		{
			Keys:    bson.M{"_id": "hashed"},
			Options: options.Index().SetBackground(true),
		},
	}
}
