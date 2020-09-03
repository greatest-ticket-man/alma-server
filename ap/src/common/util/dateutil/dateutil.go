package dateutil

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// https://pkg.go.dev/google.golang.org/protobuf@v1.25.0/types/known/timestamppb?tab=doc

// TimeToTimestamp golangのtimeからTimestampを作成する
func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// TimestampToTime Timestampからgolangのtimeに変換する
func TimestampToTime(timestamp *timestamppb.Timestamp) time.Time {
	return timestamp.AsTime()
}
