package helloctrl

import "net/http"

// PageHTML .
func PageHTML(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("Hello World"))
}
