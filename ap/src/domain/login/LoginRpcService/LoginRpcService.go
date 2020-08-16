package LoginRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/infrastructure/grpc/proto/login"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"time"
)

// Login ログイン処理
func Login(ctx context.Context, txTime time.Time, email string, password string) *login.LoginReply {

	// 取得と同時に、LoginTimeを更新
	userAccount := UserAccountRepository.FindAndUpdate(ctx, email, txTime)
	if userAccount == nil {
		chk.LE(errmsg.LoginFailed)
	}

	// passwordが正しいかを確認
	if userAccount.Password != password {
		chk.LE(errmsg.LoginWrongPassword)
	}

	token := jwt.New(txTime, userAccount.ID, email)

	return &login.LoginReply{
		Token: token,
		Name:  userAccount.Name,
		Email: userAccount.Email,
	}
}
