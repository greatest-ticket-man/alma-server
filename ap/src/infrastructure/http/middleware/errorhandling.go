package middleware

import (
	"alma-server/ap/src/common/error/almaerror"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/util/httputil/response"
	"log"
	"net/http"

	stripego "github.com/stripe/stripe-go/v71"
)

// errorhandling panicが発生した時にCatchするMiddleware

// ErrorHandlingMiddleware panicが発生した時にCatchするMiddleware
func ErrorHandlingMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	defer func() {
		if panicErr := recover(); panicErr != nil {
			doServerErrorProcess(w, panicErr)
		}
	}()

	next(w, r)
}

// doServerErrorProcess
// TODO Logを流すためのあれそれ
func doServerErrorProcess(w http.ResponseWriter, err interface{}) {

	var reason string
	switch e := err.(type) {

	case *almaerror.SystemError:
		reason = "エラーが発生しました" // System的なErrorなのでViewには見せない
		log.Printf("[SYSTEM ERROR] statuscode=%d msgcode=%s msg=%s err=%v", e.StatusCode, e.MessageCode,
			errmsg.Get("ja", e.MessageCode, e.Params...),
			e.Err,
		)
	case *almaerror.LogicError:
		reason = errmsg.Get("ja", e.MessageCode, e.Params...)
		log.Printf("[LOGIC ERROR] statuscode=%d msgcode=%s msg=%s", e.StatusCode, e.MessageCode, reason)
	case *almaerror.BillingError:
		// TODO
		log.Println("Billing errorです", e)
		if stripeErr, ok := err.(*stripego.Error); ok {

			switch stripeErr.Code {
			case stripego.ErrorCodeAPIKeyExpired:
				// たくさん
			}

		}

		reason = "課金Errorです"
	case error:
		log.Printf("Unknown error: %+v\n", e)
		reason = "不明なエラーが発生しました"
	default:
		log.Println("到達不能Errorです")
		reason = "到達不能Errorが発生しました"

	}

	// response
	response.ERROR(w, reason)

}
