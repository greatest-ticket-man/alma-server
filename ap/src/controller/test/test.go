package test

import (
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/event/EventService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML UIテスト画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	ctx := r.Context()

	result := EventService.GetEventName(ctx, req.Event)

	response.BaseHTML(
		w,
		"テスト",
		"",
		map[string]interface{}{},
		"/template/controller/test/test.html",
		map[string]interface{}{},
		[]string{},
		[]string{},
		result,
		nil,
	)
}
