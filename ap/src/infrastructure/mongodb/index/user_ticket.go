package index

import (
	"alma-server/ap/src/repository/user/ticket/UserTicketRepository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {

	userIndexMap[UserTicketRepository.ThisCollectionName] = []mongo.IndexModel{
		{
			Keys: bson.D{
				{Key: UserTicketRepository.FEventID, Value: 1},
				{Key: UserTicketRepository.FTicketID, Value: 1},
			},
			Options: options.Index().SetBackground(true).SetUnique(true),
		},
	}

}
