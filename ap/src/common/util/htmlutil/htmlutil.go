package htmlutil

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/infrastructure/file/almafile"
	"bytes"
	"html/template"
	"io"
	"io/ioutil"
	"path/filepath"
)

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
