package login

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/CommonHTMLService"
	"alma-server/ap/src/domain/login/LoginRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/login"
	"encoding/json"
	"net/http"
	"time"
)

// PageHTML ログイン画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/login/page.html",
		map[string]interface{}{
			"head":   CommonHTMLService.GetHead(),
			"header": CommonHTMLService.GetHeader(),
			"footer": CommonHTMLService.GetFooter(),
		},
	)
}

// Login ログイン処理
func Login(w http.ResponseWriter, r *http.Request) {

	// param
	req := &login.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	chk.SE(err)

	// time
	txTime := time.Now()

	// logic
	result := LoginRpcService.Login(r.Context(), txTime, req.Email, req.Password)

	// response
	response.JSON(w, result)
}
