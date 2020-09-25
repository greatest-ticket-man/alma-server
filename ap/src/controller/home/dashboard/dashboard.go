package dashboard

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventService"
	"net/http"
)

// PageHTML Dashboard
func PageHTML(w http.ResponseWriter, r *http.Request) {
	eventID := param.Value(r, "event")

	// なにも指定されていない場合は、Emptyにリダイレクト
	if eventID == "" {
		response.RedirectHTML(w, r, "/home/dashboard/empty")
		return
	}

	userEvent := EventService.GetEvent(r.Context(), eventID)
	if userEvent == nil {
		chk.LE(errmsg.EventNotFound)
	}

	response.BaseHTML(
		w,
		"ダッシュボード",
		"",
		nil,
		"/template/controller/home/dashboard/dashboard.html",
		nil,
		[]string{"/static/js/controller/home/dashboard/dashboard.js"},
		[]string{"/static/css/controller/home/dashboard/dashboard.css"},
		userEvent.Name,
		nil,
	)
}

// PageHTMLEmpty イベントが選択されていない場合はここにRedirect
func PageHTMLEmpty(w http.ResponseWriter, r *http.Request) {

	response.BaseHTML(
		w,
		"ダッシュボード",
		"",
		nil,
		"/template/controller/home/dashboard/empty/dashboard_empty.html",
		nil,
		[]string{"/static/js/controller/home/dashboard/empty/dashboard_empty.js"},
		[]string{"/static/css/controller/home/dashboard/empty/dashboard_empty.css"},
		"",
		nil,
	)
}
