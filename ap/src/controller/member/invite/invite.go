package invite

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/member/MemberInviteRpcService"
	"alma-server/ap/src/domain/menu/MenuService"
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
	txTime := almactx.GetTxTime(ctx)

	result := MemberInviteRpcService.PageHTML(ctx, mid, txTime, req.Event)

	response.BaseHTML(
		w,
		"招待中メンバー情報",
		"",
		map[string]interface{}{},
		"/template/controller/member/invite/member_invite.html",
		map[string]interface{}{
			"memberInviteInfoList": result.MemberInviteInfoList,
		},
		[]string{},
		[]string{
			"/static/css/common/table/table.css",
			"/static/css/controller/member/invite/member_invite.css",
		},
		result.EventName,
		MenuService.GetMenu("member_top", "member_invite"),
	)
}
