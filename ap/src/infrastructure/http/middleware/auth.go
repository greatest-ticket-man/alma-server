package middleware

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/util/cookieutil"
	"alma-server/ap/src/common/util/httputil/response"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// AuthMiddleware 認証
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// cookieからtokenを取得
	tokenStr := cookieutil.GetCookie(r, "token")

	token, err := jwt.Parse(tokenStr)
	if err != nil || !token.Valid {
		// redirect
		response.RedirectHTML(w, r, fmt.Sprintf("/login?fallback=%s", url.QueryEscape(r.RequestURI)))
		return
	}

	// commonDataを追加する
	claimMap := jwt.GetClaimMap(token)
	commonData := &almactx.CommonData{
		TxTime: time.Now(),
		Mid:    claimMap["mid"].(string),
		Email:  claimMap["email"].(string),
	}

	ctx := almactx.WithData(r.Context(), commonData)

	// 新しいtokenをセットする(refresh token)
	cookieutil.SetCookie(w, commonData.TxTime, "token", jwt.New(commonData.TxTime, commonData.Mid, commonData.Email))

	next(w, r.WithContext(ctx))
}
