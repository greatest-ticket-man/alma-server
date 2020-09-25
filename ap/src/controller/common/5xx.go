package common

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// InternalServerErrorPageHTML .
func InternalServerErrorPageHTML(w http.ResponseWriter, r *http.Request, reason string) {

	response.BaseHTML(
		w,
		"5xx Server Error",
		"",
		nil,
		"/template/common/5xx.html",
		map[string]interface{}{
			"reason": reason,
		},
		[]string{},
		[]string{},
		"",
		nil,
	)
}
