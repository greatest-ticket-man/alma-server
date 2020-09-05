package htmlutil

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/projectpathap"
	"alma-server/ap/src/common/util/dateutil"
	"bytes"
	"fmt"
	"html/template"
	"io"
	"path/filepath"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// templateに追加したい、functionをココに追加する
var funcMap = template.FuncMap{
	"add": func(i1 int, i2 int) int {
		return i1 + i2
	},
	"timestampToTime": func(ts *timestamppb.Timestamp) time.Time {
		return dateutil.TimestampToTime(ts)
	},
}

// Template .
func Template(w io.Writer, path string, data interface{}) {

	temp := getTemplateFromPath(path)

	temp.Execute(w, data)
}

// CreateTemplateToString pathからtemplateを作成する
func CreateTemplateToString(path string, data interface{}) template.HTML {

	buf := new(bytes.Buffer)

	Template(buf, path, data)

	return template.HTML(buf.String())
}

// getTemplateFromPath pathからtemplateを取得する
func getTemplateFromPath(path string) *template.Template {

	// Newにファイル名を与える必要がある
	tname := filepath.Base(path)

	temp, err := template.New(tname).Funcs(funcMap).ParseFiles(fmt.Sprintf("%s/asset%s", projectpathap.GetRoot(), path))
	chk.SE(err)
	return temp
}
