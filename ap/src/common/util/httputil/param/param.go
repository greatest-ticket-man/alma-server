package param

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Map 取得したparamをmapにして返す
func Map(r *http.Request) map[string]string {
	return mux.Vars(r)
}

// Value 指定したkeyに対するvalueおｗ取得する
func Value(r *http.Request, key string) string {
	return r.FormValue(key)
}

// JSON .
func JSON(r *http.Request, req interface{}) error {
	return json.NewDecoder(r.Body).Decode(req)
}
