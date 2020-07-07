package middleware

import (
	"alma-server/ap/src/common/error/almaerror"
	"log"
	"net/http"

	stripego "github.com/stripe/stripe-go/v71"
)

// errorhandling panicが発生した時にCatchするMiddleware

// ErrorHandlingMiddleware panicが発生した時にCatchするMiddleware
func ErrorHandlingMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	defer func() {
		if panicErr := recover(); panicErr != nil {
			doServerErrorProcess(panicErr)
		}
	}()

	next(w, r)
}

// doServerErrorProcess
func doServerErrorProcess(err interface{}) {

	switch e := err.(type) {

	case *almaerror.SystemError:
		// TODO
		log.Println("system errorです", e)
	case *almaerror.LogicError:
		// TODO
		log.Println("Logic errorです", e)
	case *almaerror.BillingError:
		// TODO
		log.Println("Billing errorです", e)
		if stripeErr, ok := err.(*stripego.Error); ok {

			switch stripeErr.Code {
			case stripego.ErrorCodeAPIKeyExpired:
				// たくさん
			}

		}

	case error:
		// TODO
		log.Println("Unknown errorです", e)
	default:
		// TODO
		log.Println("到達不能Errorです")

	}

}
