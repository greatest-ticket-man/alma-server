package SignupRpcService

import (
	"context"
	"log"
	"time"
)

// Signup サインアップ
func Signup(ctx context.Context, txTime time.Time, email string, password string) bool {

	log.Println("signup")
	log.Println("email is ", email, "pass is ", password)

	return true
}
