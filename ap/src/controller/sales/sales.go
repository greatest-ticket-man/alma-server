package sales

import (
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML 売上画面のトップ
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	_ = &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	response.BaseHTML(
		w,
		"売上状況",
		"/template/controller/sales/head.html",
		map[string]interface{}{},
		"/template/controller/sales/sales.html",
		map[string]interface{}{},
		[]string{},
		[]string{},
		"eventNameTODO",
		nil,
	)
}
