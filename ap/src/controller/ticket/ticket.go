package ticket

import (
	"alma-server/ap/src/common/util/httputil/param"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"net/http"
)

// PageHTML .
func PageHTML(w http.ResponseWriter, r *http.Request) {

	// param
	req := &common.EventRequest{
		Event: param.Value(r, "event"),
	}

	response.BaseHTML(
		w,
		"チケット",
		"",
		map[string]interface{}{},
		"/template/controller/ticket/ticket.html",
		map[string]interface{}{},
	)

}
