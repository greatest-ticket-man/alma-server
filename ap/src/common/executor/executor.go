package executor

import (
	"alma-server/ap/src/common/logger"
	"context"
	"errors"
)

// MongoDBなどトランザクションが使えないデータベースを使用している場合、
// ロールバック処理が煩雑になるため、こちらの処理を使用してロールバックを一括で管理します。
// また、panicの仕様を前提としているため、通常のif errにしたい場合は、別途作る必要があります。

// Unit .
type Unit struct {
	Execute        func() interface{}
	ExecuteResult  interface{}
	Rollback       func(context.Context) interface{}
	RollbackResult interface{}
}

// TestPanicFlg テストだけで使用するpanicエラー用のフラグ
var TestPanicFlg = false

// TestPanicMsg テストパニックが発生したときのエラー文言
var TestPanicMsg = "test executor error"

// Do 渡されたUnitを実行していく
func Do(units ...*Unit) {
	eIndex := 0 // ロールバックのときにどこまで実施したかが必要なため、ここで宣言して使う

	defer func() {

		if err := recover(); err != nil {

			// 時間切れのcontextを使うとロールバックできないので、時間指定なしのcontextを生成
			ctx := context.Background()

			// rollback
			for i := 0; i < eIndex; i++ {
				units[i].Rollback(ctx)
			}

			// ロールバック実行時にログを出力しておく
			logger.Info("executor rollback execed !!!")

			// re throw
			panic(err)
		}

	}()

	// execute
	for eIndex = 0; eIndex < len(units); eIndex++ {
		units[eIndex].Execute()
	}

	// テスト時のみ使用
	if TestPanicFlg {
		panic(errors.New(TestPanicMsg))
	}
}

// CreateDoNothingUnit 何もしないUnitを作成する
func CreateDoNothingUnit() *Unit {
	return &Unit{
		Execute:        func() interface{} { return nil },
		ExecuteResult:  nil,
		Rollback:       func(ctx context.Context) interface{} { return nil },
		RollbackResult: nil,
	}
}
