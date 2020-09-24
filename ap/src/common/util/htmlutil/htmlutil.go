package htmlutil

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/dateutil"
	"alma-server/ap/src/common/util/stringutil"
	"alma-server/ap/src/infrastructure/file/almafile"
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// templateに追加したい、functionをココに追加する
var funcMap = template.FuncMap{
	"add": func(i1 int, i2 int) int {
		return i1 + i2
	},
	"timestampToTime": func(ts *timestamppb.Timestamp) string {
		t := dateutil.TimestampToTime(ts)
		return dateutil.FormatYyyyMmDdSs(t)
	},
	"addComma": func(num int32) string {
		return stringutil.AddComma(num)
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

	assetFileSystem := almafile.GetAssetFileSystem(config.ConfigData)

	file, err := assetFileSystem.Open(path)
	defer func() {
		err := file.Close()
		chk.SE(err)
	}()
	chk.SE(err)

	content, err := ioutil.ReadAll(file)
	chk.SE(err)

	temp, err := template.New(tname).Funcs(funcMap).Parse(string(content))
	chk.SE(err)

	return temp
}
