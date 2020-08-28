package index

import (
	"alma-server/ap/src/repository/user/event/UserEventMemberRepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	userIndexMap[UserEventMemberRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			Keys:    bson.M{UserEventMemberRepository.FEventID: 1},
			Options: options.Index().SetBackground(true),
		},
		{
			Keys:    bson.M{UserEventMemberRepository.FMid: 1},
			Options: options.Index().SetBackground(true),
		},
	}

}
