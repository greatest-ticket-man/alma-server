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
		// TODO
		log.Println("system errorです", e)
		reason = "エラーが発生しました"
	case *almaerror.LogicError:
		// TODO req statuscode emsgとかをどうにかする
		log.Println("Logic errorです", e)
		reason = errmsg.Get("ja", e.MessageCode)
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
		// TODO
		log.Println("Unknown errorです", e)
		reason = "不明なエラーが発生しました"
	default:
		// TODO
		log.Println("到達不能Errorです")
		reason = "到達不能Errorが発生しました"

	}

	// response
	response.ERROR(w, reason)

}
