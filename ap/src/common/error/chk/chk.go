package chk

import (
	"alma-server/ap/src/common/error/almaerror"
	"net/http"
)

// SE SystemError: 500
func SE(err error, params ...interface{}) {
	if err != nil {
		panic(&almaerror.SystemError{Err: err, StatusCode: http.StatusInternalServerError, Params: params})
	}
}

// LE LogicError: 400
func LE(msgCode string, params ...interface{}) {
	panic(&almaerror.LogicError{StatusCode: http.StatusBadRequest, MessageCode: msgCode, Params: params})
}

// BE BillingError
func BE(err error, params ...interface{}) {
	if err != nil {
		panic(&almaerror.BillingError{Err: err, Params: params})
	}
}
