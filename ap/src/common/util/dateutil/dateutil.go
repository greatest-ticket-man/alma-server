package dateutil

import (
	"alma-server/ap/src/common/error/chk"
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

	// YyyyMmDdTHhSs Formで取得した時の形式
	YyyyMmDdTHhSs = "2006-01-02T15:04"
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

// ParseFormatStringToTime 指定したフォーマットの文字列をtime.Timeに変換する
func ParseFormatStringToTime(format string, strTime string) time.Time {
	t, err := time.Parse(format, strTime)
	chk.SE(err)
	return t
}

// ParseFormStrToTime フォームでのフォーマットを変更して、time.Timeに変換する
func ParseFormStrToTime(strTime string) time.Time {
	return ParseFormatStringToTime(YyyyMmDdTHhSs, strTime)
}
