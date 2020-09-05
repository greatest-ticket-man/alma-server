package dateutil

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// YyyyMmDd .
	YyyyMmDd = "20060102"

	// YyyyMm .
	YyyyMm = "200601"

	// YyyyMmDdSs .
	YyyyMmDdSs = "2006/01/02 15:04:05"
)

// FormatYyyyMmDdSs 日付文字列を返す (例：2019-12-07 11:34:53）
func FormatYyyyMmDdSs(t time.Time) string {
	return t.Format(YyyyMmDdSs)
}

// TimeToTimestamp golangのtimeからTimestampを作成する
func TimeToTimestamp(t time.Time) *timestamppb.Timestamp {
	return timestamppb.New(t)
}

// TimestampToTime Timestampからgolangのtimeに変換する
func TimestampToTime(timestamp *timestamppb.Timestamp) time.Time {
	return timestamp.AsTime()
}
