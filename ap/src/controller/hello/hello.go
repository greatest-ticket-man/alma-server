package hello

import (
	"alma-server/ap/src/common/util/httputil/response"
	"net/http"
)

// HTML .
func HTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(w, "/template/controller/hello/page.html", map[string]interface{}{
		"": "hoge",
	})
}

// API .
func API(w http.ResponseWriter, r *http.Request) {
	response.JSON(w, "hge")
}
