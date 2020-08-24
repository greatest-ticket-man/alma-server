package index

import (
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	userIndexMap[UserEventRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			// members.mid
			Keys:    bson.M{fmt.Sprintf("%s.%s", UserEventRepository.FMemberList, UserEventRepository.FMemberInfoMid): 1},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
		{
			// tmps.email
			Keys:    bson.M{fmt.Sprintf("%s.%s", UserEventRepository.FTempMemberList, UserEventRepository.FTempMemberInfoEmail): 1},
			Options: options.Index().SetUnique(true).SetBackground(true),
		},
	}
}
