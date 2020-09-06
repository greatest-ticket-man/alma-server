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
		"",
		nil,
		"/template/controller/event/event.html",
		map[string]interface{}{
			"eventId":   result.EventId,
			"eventName": result.EventName,
		},
		[]string{"/static/js/controller/event/event.js"},
		[]string{"/static/css/controller/event/event.css"},
		result.EventName,
		nil,
	)
}

// CreatePageHTML イベント作成画面
func CreatePageHTML(w http.ResponseWriter, r *http.Request) {

	result := EventRpcService.CreatePage()

	response.BaseHTML(
		w,
		"イベント作成",
		"",
		nil,
		"/template/controller/event/create/event_create.html",
		map[string]interface{}{
			"eventForm": htmlutil.CreateTemplateToString("/template/component/event/form.html", map[string]interface{}{
				"eventAuthInfoList": result.EventAuthInfoList,
			}),
		},
		[]string{
			"/static/js/util/validation/validation.js",
			"/static/js/component/event/form.js",
			"/static/js/controller/event/create/event_create.js",
		},
		[]string{
			"/static/css/component/event/form.css",
			"/static/css/controller/event/create/event_create.css",
		},
		"",
		nil,
	)
}

// UpdatePageHTML event update form
func UpdatePageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.UpdatePage(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"イベント情報編集",
		"",
		nil,
		"/template/controller/event/update/event_update.html",
		map[string]interface{}{
			"eventForm": htmlutil.CreateTemplateToString("/template/component/event/form.html", map[string]interface{}{
				"eventName":         result.EventName,
				"eventOrganization": result.OrganizationName,
				"eventAuthInfoList": result.EventAuthInfoList,
			}),
		},
		[]string{
			"/static/js/util/validation/validation.js",
			"/static/js/component/event/form.js",
			"/static/js/controller/event/update/event_update.js",
		},
		[]string{
			"/static/css/component/event/form.css",
			"/static/css/controller/event/update/event_update.css",
		},
		result.EventName,
		nil,
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
	reply := EventRpcService.CreateEvent(ctx, mid, txTime, req.EventName, req.OrganizationName, req.InviteMemberList)

	// response
	response.JSON(w, reply)
}

// UpdateEvent イベントの更新
func UpdateEvent(w http.ResponseWriter, r *http.Request) {

	// param
	req := &event.UpdateEventRequest{}
	err := param.JSON(r, req)
	chk.SE(err)

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	// update event
	EventRpcService.UpdateEvent(ctx, mid, txTime, req.EventId, req.EventName, req.OrganizationName, req.InviteMemberList)

	// response
	response.JSON(w, &common.Empty{})
}

// GetEventList イベントのリストを取得する
func GetEventList(w http.ResponseWriter, r *http.Request) {

	req := &event.GetEventListRequest{
		SearchText: param.Value(r, "search_text"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	result := EventRpcService.GetEventList(ctx, mid, txTime, req.SearchText)

	// response
	response.JSON(w, result)
}
