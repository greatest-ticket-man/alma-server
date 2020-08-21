package event

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
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
		"/template/controller/event/content.html",
		map[string]interface{}{
			"eventId":   result.EventId,
			"eventName": result.EventName,
		},
		"/template/controller/event/script.html",
		"/template/controller/event/css.html",
		result.EventName,
	)
}

// CreatePageHTML イベント作成画面
func CreatePageHTML(w http.ResponseWriter, r *http.Request) {

	response.BaseHTML(
		w,
		"イベント作成",
		"/template/controller/event/create/content.html",
		nil,
		"/template/controller/event/create/script.html",
		"/template/controller/event/create/css.html",
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
		"/template/controller/event/update/content.html",
		map[string]interface{}{},
		"/template/controller/event/update/script.html",
		"/template/controller/event/update/css.html",
		result.EventName,
	)
}
