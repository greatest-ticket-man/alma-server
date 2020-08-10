package cookieutil

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"net/http"
	"time"
)

const expireDuration = 10 * time.Hour

// SetCookie Cookieをセットする
func SetCookie(w http.ResponseWriter, txTime time.Time, name string, value string) bool {

	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  txTime.Add(expireDuration),
		HttpOnly: true,
		Secure:   config.ConfigData.HTTPServer.TLS, // httpsの場合だけこれを指定
	}

	http.SetCookie(w, cookie)

	return true
}

// GetCookie Cookieを取得する
func GetCookie(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	chk.SE(err)

	return cookie.Value
}
