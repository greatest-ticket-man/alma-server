package htmlutil

import (
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/common/util/stringutil"
	"html/template"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// templateに追加したい、functionをココに追加する
var funcMap = template.FuncMap{
	"add": func(i1 int, i2 int) int {
		return i1 + i2
	},
	"timestampToTime": func(ts *timestamppb.Timestamp) string {
		t := dateutil.TimestampToTime(ts)
		return dateutil.FormatYYYYMMDDhhmmss(t)
	},
	"timestampToDateTime": func(ts *timestamppb.Timestamp) string {

		// defaultの場合は空白を出力したいため
		// nilの場合は、空文字を返すようにしている
		if ts == nil {
			return ""
		}

		t := dateutil.TimestampToTime(ts)
		return dateutil.FormatYYYYMMDDhhmm(t)
	},
	"addComma": func(num int32) string {
		return stringutil.AddComma(num)
	},
}
