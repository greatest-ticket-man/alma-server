package sales

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML 売上画面のトップ
func PageHTML(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"売上状況",
		"",
		map[string]interface{}{},
		"/template/controller/sales/sales.html",
		map[string]interface{}{},
		[]string{},
		[]string{},
		result.EventName,
		MenuService.GetMenu("sales_top", "sales"),
		// nil,
		// MenuService.GetMenu("reserve_top", "reserve"),
	)
}
