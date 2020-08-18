package middleware

import (
	"alma-server/ap/src/common/almactx"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/util/cookieutil"
	"alma-server/ap/src/common/util/httputil/response"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// AuthMiddleware 認証
func AuthMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// cookieからtokenを取得
	tokenStr := cookieutil.GetCookie(r, "token")

	txTime := time.Now()

	// cookieutil.DeleteCookie(w, r, txTime, "token")

	token, err := jwt.Parse(tokenStr)
	// TODO どうにかする
	if err != nil || !token.Valid {

		log.Println("redirect...")
		log.Println("err is ", err)

		// redirect
		response.RedirectHTML(w, r, fmt.Sprintf("/login?fallback=%s", url.QueryEscape(r.RequestURI)))
		return
	}

	// commonDataを追加する
	claimMap := jwt.GetClaimMap(token)
	commonData := &almactx.CommonData{
		TxTime: txTime,
		Mid:    claimMap["mid"].(string),
		Email:  claimMap["email"].(string),
	}

	ctx := almactx.WithData(r.Context(), commonData)

	// 新しいtokenをセットする(refresh token)
	// cookieutil.SetCookie(w, commonData.TxTime, "token", jwt.New(commonData.TxTime, commonData.Mid, commonData.Email))

	next(w, r.WithContext(ctx))
}
