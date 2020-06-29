package middleware

import (
	"alma-server/ap/src/common/config"
	"log"
	"net/http"
	"strings"
)

// CorsMiddleware 違うドメインからもアクセスできるようにする
func CorsMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	config := config.ConfigData

	// CORS
	w.Header().Set("Access-Control-Allow-Origin", config.HTTPServer.AllowOrigin)

	//認証を行う
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, User-ID")
	// //必要なメソッドを許可する
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	//XMLHttpRequest Level2のpreflightをチェック
	if r.Method == "OPTIONS" {
		//ヘッダーにAuthorizationが含まれていた場合はpreflight成功
		s := r.Header.Get("Access-Control-Request-Headers")
		if strings.Contains(s, "authorization") || strings.Contains(s, "Authorization") {
			w.WriteHeader(http.StatusNoContent)
		} else {
			log.Println("Error...")
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	next(w, r)

}
