package event

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/util/htmlutil"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/common/util/jsonutil"
	"alma-server/ap/src/domain/event/EventRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"html/template"
	"log"
	"net/http"
)

// PageHTML event home
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: r.FormValue("event"),
	}

	// TODO req.EventRequest なければErrorとかメッセージを送る

	log.Println("req is ", jsonutil.Marshal(req))

	ctx := r.Context()
	mid := almactx.GetMid(ctx)

	result := EventRpcService.GetEvent(ctx, mid, req.Event)

	log.Println("result is ", result)

	response.HTML(
		w,
		"/template/component/base.html",
		map[string]interface{}{
			"mainTitle": "公演情報",
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
