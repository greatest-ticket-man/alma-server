package signup

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/signup/SignupRpcService"
	"alma-server/ap/src/infrastructure/grpc/proto/common"
	"alma-server/ap/src/infrastructure/grpc/proto/signup"
	"encoding/json"
	"net/http"
	"time"
)

// TODO 仮登録だとかそういう手順が死ぬほどあるが、
// TODO メインロジックのためにいったんスキップ

// PageHTML 新規登録画面
func PageHTML(w http.ResponseWriter, r *http.Request) {

	response.HTML(
		w,
		"/template/controller/signup/page.html",
		map[string]interface{}{},
	)

}

// Signup 新規追加
func Signup(w http.ResponseWriter, r *http.Request) {

	// param
	req := &signup.SignupRequest{}
	err := json.NewDecoder(r.Body).Decode(req)
	chk.SE(err)

	// time
	txTime := time.Now()

	// signup
	SignupRpcService.Signup(r.Context(), txTime, req.Email, req.Password)

	// response
	response.JSON(w, &common.Empty{})
}
