package ticketctrl

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// CreatePage チケット作成画面
func CreatePage(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"チケット作成V2",
		"/template/controller/ticket/create/head.html",
		response.M{},
		"/template/controller/ticket/create/ticketv2.html",
		response.M{},
		[]string{},
		[]string{
			"/static/css/common/content_head_button/content_head_button.css",
		},
		result.EventName,
		MenuService.GetMenu("ticket_top", ""),
	)

}
