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

	// TodoDeleteFailed todoの削除に失敗しました
	TodoDeleteFailed = "TodoDeleteFailed"

	// LoginFailed ログインに失敗しました
	LoginFailed = "LoginFailed"

	// LoginWrongPassword パスワードが違います
	LoginWrongPassword = "LoginWrongPassword"
)

var msgMap = map[string]map[string]string{
	"ja": {
		HelloLogicError:    "テスト用のError",
		TodoDeleteFailed:   "todoの削除に失敗しました",
		LoginFailed:        "ログインに失敗しました",
		LoginWrongPassword: "パスワードが違います",
	},
}
