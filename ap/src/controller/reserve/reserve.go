package reserve

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
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
			"/static/js/controller/reserve/reserve.js",
		},
		[]string{
			"/static/css/common/table/table.css",
			"/static/css/common/content_head_button/content_head_button.css",
		},
		result.EventName,
		MenuService.GetMenu("reserve_top", "reserve"),
	)

}

// CreatePageHTML 予約の作成画面
func CreatePageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"予約の作成",
		"/template/controller/reserve/create/head.html",
		map[string]interface{}{},
		"/template/controller/reserve/create/reserve_create.html",
		map[string]interface{}{
			"reserveForm": htmlutil.CreateTemplateToString(
				"/template/component/reserve/form.html",
				map[string]interface{}{},
			),
			"orderForm": htmlutil.CreateTemplateToString(
				"/template/component/reserve/order_form.html",
				map[string]interface{}{},
			),
		},
		[]string{
			"/static/js/component/reserve/form.js",
			"/static/js/controller/reserve/create/reserve_create.js",
		},
		[]string{
			"/static/css/component/reserve/form.css",
			// "/static/css/common/content_head_button/content_head_button.css",
			"/static/css/controller/reserve/create/head.css",
			"/static/css/controller/reserve/create/reserve_create.css",
		},
		result.EventName,
		MenuService.GetMenu("reserve_top", ""),
	)
}
