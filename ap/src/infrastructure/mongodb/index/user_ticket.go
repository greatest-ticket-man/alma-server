package index

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	userIndexMap["USER_TICKET"] = []mongo.IndexModel{
		{
			Keys:    bson.M{"_id": "hashed"},
			Options: options.Index().SetBackground(true),
		},
	}
}
