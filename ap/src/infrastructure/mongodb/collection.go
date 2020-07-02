package mongodb

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/reflectutil"
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// AlmaCollection MongoCollectionのWrapper
type AlmaCollection struct {
	col *mongo.Collection
	ctx context.Context
}

// Find .
func (c *AlmaCollection) Find(query interface{}, reflectType reflect.Type, opts ...*options.FindOptions) interface{} {
	reflectSlice := reflectutil.CreateSlice(reflectType)
	return reflectionFind(c.ctx, c.col, query, reflectSlice, reflectType, opts...)
}

// FindBsonDList .
func (c *AlmaCollection) FindBsonDList(query interface{}, opts ...*options.FindOptions) interface{} {
	reflectType := reflect.TypeOf([]primitive.E{})
	reflectSlice := reflectutil.CreateSlice(reflectType)
	return reflectionFind(c.ctx, c.col, query, reflectSlice, reflectType, opts...)
}

// GetMongoDriverCollection テスト用のメソッド。本番ではWrapperメソッドを作成して使ってください。
func (c *AlmaCollection) GetMongoDriverCollection() *mongo.Collection {
	return c.col
}

// FindOne .
func (c *AlmaCollection) FindOne(query interface{}, reflectType reflect.Type, opts ...*options.FindOneOptions) interface{} {
	result := c.col.FindOne(c.ctx, query, opts...)
	return checkAndReflectSingleResult(result, reflectType)
}

// Count .
func (c *AlmaCollection) Count(query interface{}, opts ...*options.CountOptions) int64 {
	cnt, err := c.col.CountDocuments(c.ctx, query, opts...)
	chk.SE(err)
	return cnt
}

// InsertMany .
func (c *AlmaCollection) InsertMany(list []interface{}, opts ...*options.InsertManyOptions) []interface{} {
	res, err := c.col.InsertMany(c.ctx, list, opts...)
	chk.SE(err)
	return res.InsertedIDs
}

// InsertOne .
func (c *AlmaCollection) InsertOne(data interface{}) interface{} {
	res, err := c.col.InsertOne(c.ctx, data)
	chk.SE(err)
	return res.InsertedID
}

// ReplaceOne .
func (c *AlmaCollection) ReplaceOne(query map[string]interface{}, data interface{}) int32 {
	res, err := c.col.ReplaceOne(c.ctx, query, data)
	chk.SE(err)
	return int32(res.ModifiedCount + res.UpsertedCount)
}

// RemoveAllTestOnly .
// TODO テストの時以外には簡単に使えないようにしたい。
func (c *AlmaCollection) RemoveAllTestOnly() int {
	result, err := c.col.DeleteMany(c.ctx, &map[string]interface{}{})
	chk.SE(err)
	return int(result.DeletedCount)
}

// DeleteOne .
func (c *AlmaCollection) DeleteOne(query interface{}) int32 {
	result, err := c.col.DeleteOne(c.ctx, query)
	chk.SE(err)
	return int32(result.DeletedCount)
}

// DeleteMany .
func (c *AlmaCollection) DeleteMany(query interface{}) int32 {
	result, err := c.col.DeleteMany(c.ctx, query)
	chk.SE(err)
	return int32(result.DeletedCount)
}

// UpdateOne .
func (c *AlmaCollection) UpdateOne(query interface{}, update interface{}, opts ...*options.UpdateOptions) int32 {
	res, err := c.col.UpdateOne(c.ctx, query, update, opts...)
	chk.SE(err)

	return int32(res.UpsertedCount + res.ModifiedCount)
}

// UpdateMany .
func (c *AlmaCollection) UpdateMany(query interface{}, update interface{}, opts ...*options.UpdateOptions) int32 {
	res, err := c.col.UpdateMany(c.ctx, query, update, opts...)
	chk.SE(err)

	return int32(res.UpsertedCount + res.ModifiedCount)
}

// UpsertOne .
func (c *AlmaCollection) UpsertOne(query interface{}, Upsert interface{}, opts ...*options.UpdateOptions) int32 {
	addedOpts := append(opts, options.Update().SetUpsert(true))
	return c.UpdateOne(query, Upsert, addedOpts...)
}

// BulkWrite .
func (c *AlmaCollection) BulkWrite(models []mongo.WriteModel, opts ...*options.BulkWriteOptions) int32 {

	res, err := c.col.BulkWrite(c.ctx, models, opts...)
	chk.SE(err)
	return int32(res.UpsertedCount + res.ModifiedCount)
}

// FindOneAndUpdate .
func (c *AlmaCollection) FindOneAndUpdate(query interface{}, update interface{}, reflectType reflect.Type, opts ...*options.FindOneAndUpdateOptions) interface{} {
	res := c.col.FindOneAndUpdate(c.ctx, query, update, opts...)
	return checkAndReflectSingleResult(res, reflectType)
}

// FindOneAndUpsert upsertの関係上、必ず実行後のデータを返すようにしている
func (c *AlmaCollection) FindOneAndUpsert(query interface{}, update interface{}, reflectType reflect.Type, opts ...*options.FindOneAndUpdateOptions) interface{} {
	addedOpts := append(opts, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	res := c.col.FindOneAndUpdate(c.ctx, query, update, addedOpts...)
	return checkAndReflectSingleResult(res, reflectType)
}

// FindOneAndDelete .
func (c *AlmaCollection) FindOneAndDelete(query interface{}, reflectType reflect.Type, opts ...*options.FindOneAndDeleteOptions) interface{} {
	res := c.col.FindOneAndDelete(c.ctx, query, opts...)
	return checkAndReflectSingleResult(res, reflectType)
}

// Aggregate Aggregate処理を行う
func (c *AlmaCollection) Aggregate(ctx context.Context, query interface{}, reflectType reflect.Type, opts ...*options.AggregateOptions) interface{} {
	cursor, err := c.col.Aggregate(ctx, query, opts...)
	chk.SE(err)

	reflectSlice := reflectutil.CreateSlice(reflectType)
	for cursor.Next(ctx) {
		v := reflect.New(reflectType)
		chk.SE(cursor.Decode(v.Interface()))
		reflectSlice.Set(reflect.Append(reflectSlice, v.Elem()))
	}
	chk.SE(cursor.Err())

	return reflectSlice.Interface()
}

// CreateIndex indexを作成する
func (c *AlmaCollection) CreateIndex(ctx context.Context, indexList []mongo.IndexModel) []string {
	names, err := c.col.Indexes().CreateMany(ctx, indexList)
	chk.SE(err)
	return names
}

func checkAndReflectSingleResult(result *mongo.SingleResult, reflectType reflect.Type) interface{} {
	// 検索がヒットしない場合、エラーとなる。その場合、nilを返す。
	if result.Err() != nil && result.Err().Error() == "mongo: no documents in result" {
		return nil // TODO ↑のエラーメッセージが変更される危険性があるため、テストを書いておいたほうが良いかも。
	}
	chk.SE(result.Err())
	v := reflect.New(reflectType)
	chk.SE(result.Decode(v.Interface()))
	return v.Elem().Interface()
}

// リフレクションタイプでデータを取得する
func reflectionFind(ctx context.Context, collection *mongo.Collection, query interface{}, reflectSlice reflect.Value, reflectType reflect.Type, opts ...*options.FindOptions) interface{} {
	cur, err := collection.Find(ctx, query, opts...)
	defer func() {
		if cur != nil {
			cur.Close(ctx)
		}
		if err2 := recover(); err2 != nil {
			panic(err2)
		}
	}()
	chk.SE(err)

	for cur.Next(ctx) {
		v := reflect.New(reflectType)
		chk.SE(cur.Decode(v.Interface()))
		reflectSlice.Set(reflect.Append(reflectSlice, v.Elem()))
	}
	chk.SE(cur.Err())
	return reflectSlice.Interface()
}
