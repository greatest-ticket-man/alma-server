package middleware

import (
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/util/cookieutil"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/common/util/jsonutil"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// AuthMiddleware 認証
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	log.Println("============== auth !!! ================")

	tokenStr := cookieutil.GetCookie(r, "token")

	token, err := jwt.Parse(tokenStr)
	if err != nil || !token.Valid {
		// redirect
		response.RedirectHTML(w, r, fmt.Sprintf("/login?fallback=%s", url.QueryEscape(r.RequestURI)))
		return
	}

	// TODO contextにデータを突っ込む, midとかemailとか
	log.Println("token is ", jsonutil.Marshal(token))

	next(w, r)

}
