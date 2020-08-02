package login

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/domain/CommonHTMLService"
	"encoding/json"
	"log"
	"net/http"
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
	data := &struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}{}

	err := json.NewDecoder(r.Body).Decode(data)
	chk.SE(err)

	// TODO ログイン失敗したときはError

	log.Println("name is ", data.Name, "pass is ", data.Pass)

	response.JSON(w, true)
}
