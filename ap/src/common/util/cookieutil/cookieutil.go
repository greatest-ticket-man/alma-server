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

	// cookieがない場合はから文字を送信する
	if err == http.ErrNoCookie {
		return ""
	}

	chk.SE(err)

	return cookie.Value
}

// DeleteCookie Cookieを削除する
func DeleteCookie(w http.ResponseWriter, r *http.Request, txTime time.Time, name string) {

	cookie, err := r.Cookie(name)
	if err == http.ErrNoCookie {
		return
	}
	chk.SE(err)

	// 削除
	cookie.MaxAge = -1
	cookie.Expires = txTime
	http.SetCookie(w, cookie)
}
