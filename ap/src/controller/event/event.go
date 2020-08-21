package event

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"encoding/json"
	"net/http"
)

// PageHTML event home
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"イベント情報",
		"/template/controller/event/event.html",
		map[string]interface{}{
			"eventId":   result.EventId,
			"eventName": result.EventName,
		},
		[]string{"/static/js/controller/event/event.js"},
		[]string{"/static/css/controller/event/event.css"},
		result.EventName,
	)
}

// CreatePageHTML イベント作成画面
func CreatePageHTML(w http.ResponseWriter, r *http.Request) {

	result := EventRpcService.CreatePage()

	response.BaseHTML(
		w,
		"イベント作成",
		"/template/controller/event/create/event_create.html",
		map[string]interface{}{
			"eventForm": htmlutil.CreateTemplateToString("/template/component/event/form.html", map[string]interface{}{
				"result": result,
			}),
		},
		[]string{
			"/static/js/controller/event/create/event_create.js",
			"/static/js/component/event/form.js",
		},
		[]string{
			"/static/css/component/event/form.css",
			"/static/css/controller/event/create/event_create.css",
		},
		"",
	)
}

// CreateEvent イベントの作成
func CreateEvent(w http.ResponseWriter, r *http.Request) {

	// param
	req := &event.CreateEventRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	chk.SE(err)

	ctx := r.Context()

	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	// create event
	reply := EventRpcService.CreateEvent(ctx, mid, txTime, req.EventName, req.OrganizationName, req.MemberInfoList)

	// response
	response.JSON(w, reply)
}

// UpdatePageHTML event update form
func UpdatePageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"イベント情報編集",
		"/template/controller/event/update/event_update.html",
		map[string]interface{}{
			"eventForm": htmlutil.CreateTemplateToString("/template/component/event/form.html", map[string]interface{}{
				"eventName":         result.EventName,
				"eventOrganization": "組織名",
			}),
		},
		[]string{
			"/static/js/component/event/form.js",
		},
		[]string{
			"/static/css/component/event/form.css",
			"/static/css/controller/event/update/event_update.css",
		},
		result.EventName,
	)
}
