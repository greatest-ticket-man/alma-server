package dashboard

import (
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"html/template"
	"net/http"
)

// PageHTML Dashboard
func PageHTML(w http.ResponseWriter, r *http.Request) {
	eventID := param.Value(r, "event")

	// なにも指定されていない場合は、Emptyにリダイレクト
	if eventID == "" {
		response.RedirectHTML(w, r, "/home/dashboard/empty")
	}

	userEvent := EventService.GetEvent(r.Context(), eventID)
	if userEvent == nil {
		// TODO error message
		userEvent = &UserEventRepository.UserEvent{Name: "TODO このパターンはエラーメッセージを出すようにする"}
	}

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "ダッシュボード",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/content.html", "")),
			"script":      template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/script.html", "")),
			"css":         template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/css.html", "")),
			"eventName":   userEvent.Name,
		},
	)
}

// PageHTMLEmpty イベントが選択されていない場合はここにRedirect
func PageHTMLEmpty(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle":   "ダッシュボード",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/empty/content.html", "")),
			"script":      template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/empty/script.html", "")),
			"css":         template.HTML(htmlutil.CreateTemplateToString("/template/controller/home/dashboard/empty/css.html", "")),
			"eventName":   EventService.DefaultEventName,
		},
	)
}
