package middleware

import (
	"alma-server/ap/src/common/error/almaerror"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/logger"
	"alma-server/ap/src/common/util/httputil/response"
	"alma-server/ap/src/controller/common"
	"net/http"
	"runtime"

	stripego "github.com/stripe/stripe-go/v71"
)

// ErrorHandlingMiddleware panicが発生した時にCatchするMiddleware
func ErrorHandlingMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	defer func() {
		if panicErr := recover(); panicErr != nil {

			stackTrace()

			reason := logAndReason(w, panicErr, r.URL.String())

			if r.Method == "GET" { // GET requestの場合はUIを表示する
				common.InternalServerErrorPageHTML(w, r, reason)
			} else {
				response.ERROR(w, reason)
			}
		}
	}()

	next(w, r)
}

// stackTrace Errorが発生下場所を表示する
func stackTrace() {
	stack := make([]byte, 1024*8)
	stack = stack[:runtime.Stack(stack, false)]
	logger.Infof(string(stack))
}

func logAndReason(w http.ResponseWriter, err interface{}, url string) string {

	var reason string
	switch e := err.(type) {

	case *almaerror.SystemError:
		reason = "エラーが発生しました" // System的なErrorなのでViewには見せない

		logger.Infof("[SYSTEM ERROR] url=%s statuscode=%d msgcode=%s msg=%s err=%v", url, e.StatusCode, e.MessageCode,
			errmsg.Get("ja", e.MessageCode, e.Params...),
			e.Err,
		)
	case *almaerror.LogicError:
		reason = errmsg.Get("ja", e.MessageCode, e.Params...)
		logger.Infof("[LOGIC ERROR] url=%s statuscode=%d msgcode=%s msg=%s", url, e.StatusCode, e.MessageCode, reason)
	case *almaerror.BillingError:
		// TODO
		// log.Println("Billing errorです", e)
		if stripeErr, ok := err.(*stripego.Error); ok {

			switch stripeErr.Code {
			case stripego.ErrorCodeAPIKeyExpired:
				// たくさん
			}
		}

		reason = "課金Errorです"
	case error:
		logger.Infof("Unknown error: %+v\n", e)
		reason = "不明なエラーが発生しました"
	default:
		logger.Infof("到達不能Errorです")
		reason = "到達不能Errorが発生しました"
	}

	return reason
}
