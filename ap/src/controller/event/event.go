package event

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"encoding/json"
	"html/template"
	"net/http"
)

// PageHTML event home
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: r.FormValue("event"),
	}

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle": "イベント情報",
			"mainContent": template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/content.html",
				map[string]interface{}{
					"eventId":   result.EventId,
					"eventName": result.EventName,
				})),
			"script": template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/javascript.html", "")),
			"css":    template.HTML(htmlutil.CreateTemplateToString("/template/controller/event/css.html", "")),
		},
	)
}

// GetEvent イベント名を取得する
func GetEvent(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	chk.SE(err)

	ctx := r.Context()

	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)
	response.JSON(w, result)
}
