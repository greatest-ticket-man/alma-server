package reserve

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/domain/reserve/ReserveRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"alma-server/ap/src/infrastructure/grpc/proto/reserve"
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
	txTime := almactx.GetTxTime(ctx)

	result := ReserveRpcService.CreatePage(ctx, mid, txTime, req.Event)

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
				map[string]interface{}{
					// "calendar": htmlutil.CreateTemplateToString(
					// 	"/template/component/ui/calendar/calendar2.html",
					// 	map[string]interface{}{},
					// ),
					"ticketInfoList": result.TicketInfoList,
				},
			),
			"customerForm": htmlutil.CreateTemplateToString(
				"/template/component/reserve/customer_form.html",
				map[string]interface{}{},
			),
			"payForm": htmlutil.CreateTemplateToString(
				"/template/component/reserve/pay_form.html",
				map[string]interface{}{},
			),
		},
		[]string{
			// "/static/js/component/ui/calendar/calendar2.js", // calendar
			"/static/js/component/reserve/order_form.js",
			"/static/js/component/reserve/form.js",
			"/static/js/controller/reserve/create/reserve_create.js",
		},
		[]string{
			// "/static/css/component/ui/calendar/calendar2.css", // calendar
			"/static/css/component/reserve/form.css",
			"/static/css/controller/reserve/create/head.css",
			"/static/css/controller/reserve/create/reserve_create.css",
		},
		result.EventName,
		MenuService.GetMenu("reserve_top", ""),
	)
}

// CreateReserve 予約の作成
func CreateReserve(w http.ResponseWriter, r *http.Request) {

	// param
	req := &reserve.CreateReserveRequest{}
	err := param.JSON(r, req)
	chk.SE(err)

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	ReserveRpcService.CreateReserve(
		ctx, mid, txTime,
		req.EventId,
		req.TicketId,
		req.ScheduleId,
		req.TicketNum,
		req.Desc,
		req.Name,
		req.NameFurigana,
		req.Email,
		req.PayKind,
	)

	response.JSON(w, &common.Empty{})
}
