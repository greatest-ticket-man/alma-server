package UserSequenceRepository

import (
	"alma-server/ap/src/infrastructure/mongodb"
	"context"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// ThisCollectionName .
	ThisCollectionName = "USER_SEQUENCE"

	// field
	fid = "_id"

	// FEid .
	FEid = "eid"

	// FKey .
	FKey = "key"
	fseq = "seq"

	// ReserveKey 予約番号
	ReserveKey seqKey = "reserve"
)

var reflectType = reflect.TypeOf((*UserSequence)(nil))

type seqKey string

// UserSequence .
type UserSequence struct {
	ID      *primitive.ObjectID `bson:"_id,omitempty"`
	EventID string              `bson:"eid"`
	Key     string              `bson:"key"`
	Seq     uint64              `bson:"seq"`
}

func getDb(ctx context.Context) *mongodb.AlmaCollection {
	return mongodb.GetUserCollection(ctx, ThisCollectionName)
}

// Next .
func Next(ctx context.Context, eventID string, key seqKey) uint64 {
	return Increment(ctx, eventID, key, 1)
}

// Increment .
func Increment(ctx context.Context, eventID string, key seqKey, count uint64) uint64 {

	query := bson.M{FEid: eventID, FKey: key}

	upsert := bson.M{
		"$inc": bson.M{
			fseq: count,
		},
		"$setOnInsert": bson.M{
			FEid: eventID,
			FKey: key,
		},
	}

	opt := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(true)

	result := getDb(ctx).FindOneAndUpdate(query, upsert, reflectType, opt)
	if result == nil {
		return 0
	}

	return result.(*UserSequence).Seq
}
