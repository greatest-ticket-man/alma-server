package errmsg

import "fmt"

// Get .
func Get(locale string, msgCode string, params ...interface{}) string {
	msg := msgMap[locale][msgCode]
	return fmt.Sprintf(msg, params...)
}

const (
	// HelloLogicError テスト用のError
	HelloLogicError = "HelloLogicError"
)

var msgMap = map[string]map[string]string{
	"ja": {
		HelloLogicError: "テスト用のError",
	},
}
