package response

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/infrastructure/grpc/proto/menu"
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
func BaseHTML(w http.ResponseWriter, mainTitle string, headContentPath string, headContentpath map[string]interface{}, mainContentPath string, mainDataMap map[string]interface{},
	scriptPathList []string, cssPathList []string, eventName string, mstMenu *menu.MenuInfo) {

	if eventName == "" {
		eventName = "イベントの選択"
	}

	// headContentはpathがある場合のみ追加する
	headContent := template.HTML("")
	if headContentPath != "" {
		headContent = template.HTML(htmlutil.CreateTemplateToString(headContentPath, headContentpath))
	}

	// sidebar sidebarの指定がある場合のみ、SideBarを読み込む
	var sideBar template.HTML
	if mstMenu != nil {
		sideBar = template.HTML(htmlutil.CreateTemplateToString("/template/component/base/side_bar.html", mstMenu))
	}

	// TODO sidemenu

	HTML(
		w,
		"/template/component/base/base.html",
		map[string]interface{}{
			"mainTitle":      mainTitle,
			"headContent":    headContent,
			"mainContent":    template.HTML(htmlutil.CreateTemplateToString(mainContentPath, mainDataMap)),
			"scriptPathList": scriptPathList,
			"cssPathList":    cssPathList,
			"eventName":      eventName,
			"sideBar":        sideBar,
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
