package member

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/member/MemberRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML member home
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)
	txTime := almactx.GetTxTime(ctx)

	result := MemberRpcService.PageHTML(ctx, mid, txTime, req.Event)

	response.BaseHTML(
		w,
		"メンバー情報",
		"/template/controller/member/member.html",
		map[string]interface{}{},
		nil,
		nil,
		result.EventName,
	)

}
