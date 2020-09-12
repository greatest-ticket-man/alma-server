package ticket

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML .
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
		"チケット",
		"/template/controller/ticket/head.html",
		map[string]interface{}{},
		"/template/controller/ticket/ticket.html",
		map[string]interface{}{},
		[]string{
			"/static/js/controller/ticket/ticket.js",
		},
		[]string{
			"/static/css/component/common/content_head_button/content_head_button.css",
			"/static/css/controller/ticket/ticket.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", "ticket"),
	)
}

// CreatePageHTML チケットの作成画面
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
		"チケット作成",
		"/template/controller/ticket/create/head.html",
		map[string]interface{}{},
		"/template/controller/ticket/create/ticket_create.html",
		map[string]interface{}{},
		[]string{},
		[]string{
			"/static/css/component/common/content_head_button/content_head_button.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", ""),
	)
}
