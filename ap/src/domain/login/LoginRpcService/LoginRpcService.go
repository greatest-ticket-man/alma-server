package LoginRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/infrastructure/grpc/proto/login"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"time"
)

// Login ログイン処理
func Login(ctx context.Context, txTime time.Time, email string, password string) *login.LoginReply {

	// TODO 同時にログイン時間を更新したい
	userAccount := UserAccountRepository.GetFromEmail(ctx, email)
	if userAccount == nil {
		chk.LE(errmsg.LoginFailed)
	}

	// passwordが正しいかを確認
	if userAccount.Password != password {
		chk.LE(errmsg.LoginWrongPassword)
	}

	// TODO 正しければJWT tokenを生成する
	return &login.LoginReply{
		Token: "hoge",
	}
}
