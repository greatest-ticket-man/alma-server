package invite

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML member invite home
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	// txTime := almactx.GetTxTime(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.BaseHTML(
		w,
		"招待中メンバー情報",
		"",
		map[string]interface{}{},
		"/template/controller/member/invite/member_invite.html",
		map[string]interface{}{},
		[]string{},
		[]string{
			"/static/css/controller/member/invite/member_invite.css",
		},
		result.EventName,
	)

}
