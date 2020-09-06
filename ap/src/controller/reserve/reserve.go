package reserve

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
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

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"予約状況",
		"",
		map[string]interface{}{},
		"/template/controller/reserve/reserve.html",
		map[string]interface{}{},
		[]string{},
		[]string{},
		result.EventName,
		MenuService.GetMenu("reserve_top", "reserve"),
	)

}
