package index

import (
	"alma-server/ap/src/repository/user/UserSequenceRepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	userIndexMap[UserSequenceRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: UserSequenceRepository.FEid, Value: 1},
				{Key: UserSequenceRepository.FKey, Value: 1},
			},
			Options: options.Index().SetUnique(true),
		},
	}
}
