package event

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
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
		"/template/controller/event/javascript.html",
		"/template/controller/event/css.html",
		result.EventName,
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
