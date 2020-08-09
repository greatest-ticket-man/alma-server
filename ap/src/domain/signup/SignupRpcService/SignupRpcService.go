package SignupRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/util/uniqueidutil"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"time"
)

// Signup サインアップ
func Signup(ctx context.Context, txTime time.Time, email string, password string) bool {

	// 先に、指定されたemailがすでに存在するかどうかを確認する
	userAccount := UserAccountRepository.GetFromEmail(ctx, email)
	if userAccount != nil {
		// このメールアドレスはすでに使用されています
		chk.LE(errmsg.SignupAlreadyUseEmail)
	}

	// 作成
	UserAccountRepository.Insert(ctx,
		txTime,
		uniqueidutil.GenerateUniqueID(),
		email,
		password,
	)
	return true
}
