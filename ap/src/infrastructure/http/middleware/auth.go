package middleware

import (
	"log"
	"net/http"
)

// AuthMiddleware 認証
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// TODO pathがstaticなら、スキップ

	// token := jwt.Auth(r)

	// もし、認証が合わないとか、有効期限が切れたとかなったら、callbackのパスをつけてlogin画面に遷移させる

	// TODO 認証がとおった場合は新しいTokenを送信する

	log.Println("============== auth !!! ================")

	next(w, r)

}
