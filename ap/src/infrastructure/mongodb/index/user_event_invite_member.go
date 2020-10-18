package index

import (
	"alma-server/ap/src/repository/user/event/UserEventInviteMemberRepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	userIndexMap[UserEventInviteMemberRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			Keys:    bson.M{UserEventInviteMemberRepository.FEventID: 1},
			Options: options.Index(),
		},
	}
}
