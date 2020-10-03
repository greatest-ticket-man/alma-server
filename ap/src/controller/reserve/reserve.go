package reserve

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/domain/reserve/ReserveRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML reserve page
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	result := ReserveRpcService.Page(ctx, mid, txTime, req.Event)

	response.BaseHTML(
		w,
		"予約状況",
		"/template/controller/reserve/head.html",
		map[string]interface{}{},
		"/template/controller/reserve/reserve.html",
		map[string]interface{}{
			"reserveInfoList": result.ReserveInfoList,
			"eventId":         result.EventId,
		},
		[]string{
			"/static/js/common/table/table.js",
		},
		[]string{
			"/static/css/common/table/table.css",
			"/static/css/common/content_head_button/content_head_button.css",
		},
		result.EventName,
		MenuService.GetMenu("reserve_top", "reserve"),
	)

}
