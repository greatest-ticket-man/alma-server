package htmlutil

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/projectpathap"
	"bytes"
	"fmt"
	"html/template"
	"io"
)

// Template .
func Template(w io.Writer, path string, data interface{}) {

	temp := getTemplateFromPath(path)

	temp.Execute(w, data)
}

// CreateTemplateToString pathからtemplateを作成する
func CreateTemplateToString(path string, data interface{}) string {

	buf := new(bytes.Buffer)

	Template(buf, path, data)

	return buf.String()
}

// getTemplateFromPath pathからtemplateを取得する
func getTemplateFromPath(path string) *template.Template {
	temp, err := template.ParseFiles(fmt.Sprintf("%s/asset%s", projectpathap.GetRoot(), path))
	chk.SE(err)
	return temp
}
