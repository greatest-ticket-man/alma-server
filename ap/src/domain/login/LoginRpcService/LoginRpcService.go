package LoginRpcService

import (
	"alma-server/ap/src/infrastructure/grpc/proto/login"
	"context"
	"log"
	"time"
)

// Login ログイン処理
func Login(ctx context.Context, txTime time.Time, email string, password string) *login.LoginReply {

	log.Println("email is ", email, "pass is ", password)

	return &login.LoginReply{
		Token: "hoge",
	}
}
