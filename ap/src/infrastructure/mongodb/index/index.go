package index

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var userIndexMap = map[string][]mongo.IndexModel{}

// CreateIndex MongoDBのindex作成
func CreateIndex(ctx context.Context) {

	// userdb index
	createDbIndex(ctx, mongodb.UserDbName, userIndexMap)
}

func createDbIndex(ctx context.Context, dbName string, indexMap map[string][]mongo.IndexModel) {
	for colName, indexList := range indexMap {
		col := mongodb.GetDbCollection(ctx, dbName, colName)
		col.CreateIndex(ctx, indexList)
	}
}
