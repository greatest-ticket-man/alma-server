package create

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/event"
	"encoding/json"
	"net/http"
)

// PageHTML イベント作成画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

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

// Event イベントの作成
func Event(w http.ResponseWriter, r *http.Request) {

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
