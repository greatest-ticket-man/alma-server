package response

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/jsonutil"
	"html/template"
	"net/http"
)

// JSON responseをjsonにする
func JSON(w http.ResponseWriter, result interface{}) {
	resultMap := map[string]interface{}{
		"result":  result,
		"success": true,
	}
	w.WriteHeader(http.StatusOK)
	jsonutil.Write(w, resultMap)
}

// HTML responseをhtmlで返す
func HTML(w http.ResponseWriter, path string, data map[string]interface{}) {
	w.WriteHeader(http.StatusOK)
	htmlutil.Template(w, path, data)
}

// BaseHTML baseテンプレートでhtmlを返す
func BaseHTML(w http.ResponseWriter, mainTitle string, mainContentPath string, mainDataMap map[string]interface{},
	scriptPath string, cssPath string, eventName string) {

	if eventName == "" {
		eventName = "イベントの選択"
	}

	HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   mainTitle,
			"mainContent": template.HTML(htmlutil.CreateTemplateToString(mainContentPath, mainDataMap)),
			"script":      template.HTML(htmlutil.CreateTemplateToString(scriptPath, "")),
			"css":         template.HTML(htmlutil.CreateTemplateToString(cssPath, "")),
			"eventName":   eventName,
		},
	)

}

// RedirectHTML redirect
func RedirectHTML(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, http.StatusFound)
}

// ERROR errorを添えて返す
// StatusOKで応答しないと、errorを出してしまうためStatusOKで返しています
func ERROR(w http.ResponseWriter, reason string) {
	resultMap := map[string]interface{}{
		"success": false,
		"reason":  reason,
	}

	w.WriteHeader(http.StatusOK)
	jsonutil.Write(w, resultMap)
}
