package LoginRpcService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/infrastructure/grpc/proto/login"
	"alma-server/ap/src/repository/user/UserAccountRepository"
	"context"
	"log"
	"time"
)

// Login ログイン処理
func Login(ctx context.Context, txTime time.Time, email string, password string) *login.LoginReply {

	log.Println("email is ", email, "pass is ", password)

	// TODO data取得
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
