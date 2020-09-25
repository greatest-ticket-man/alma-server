package common

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// NotFoundPageHTML .
func NotFoundPageHTML(w http.ResponseWriter, r *http.Request) {

	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"404 Page Not Found",
		"",
		nil,
		"/template/common/404.html",
		map[string]interface{}{},
		[]string{
			"/static/js/common/404.js",
		},
		[]string{
			"/static/css/common/404.css",
		},
		result.EventName,
		nil,
	)

}
