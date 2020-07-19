// Package jobrunner 定期的にJobを実行してくれる物です
// 実行したいJobを 「func AddJob()」 で追加することで、1秒以内に実行してくれます
// Responseで待つ必要がなく、重い処理がある場合は、こちらを使ってみてください
//
// 次の周期の実行は、前の処理が処理中の場合はスキップされます
//
// Job.Dataに格納したデータは、実行時にExecFuncの「data interface{}」という引数で渡されます
//
package jobrunner

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/almaerror"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/error/errmsg"
	"alma-server/ap/src/common/util/stringutil"
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// Jobのタイムアウト時間
const timeout = 15 * time.Second

// jobSeq LocalCacheでジョブを管理するための、sequence
// 直接参照しないでください、getJobSeqから取得してください
var jobSeq uint32

// createJobSeq 並行でアクセスしても、必ずuniqueな数字を得ることができる
func createJobSeq() string {
	return stringutil.Uint32ToString(atomic.AddUint32(&jobSeq, 1))
}

// runFlag Jobが走っているかを判定する
var runFlag bool = false

// Job JobのStruct
type Job struct {
	// Data この値は、ExecFuncの引数dataに渡され覚ます
	Data     interface{}
	ExecFunc func(ctx context.Context, data interface{}) interface{}
}

// AddJob 1秒周期のJobリストの追加する
func AddJob(job *Job) bool {

	if !runFlag {
		chk.SE(errors.New("jobrunnerが起動していません"))
	}

	setCacheJob(createJobSeq(), job)
	return true
}

// Run jobrunnerを起動します
func Run() {
	log.Println("Alma job runner Start")

	// set up
	cacheSetUp()

	// job run
	go func() {

		// job run flag
		runFlag = true
		defer func() {
			runFlag = false
		}()

		// ticker
		ticker := time.NewTicker(1 * time.Second)
		for {
			<-ticker.C
			doJob()
		}
	}()
}

// doJob jobMapに蓄積されたJobを実行します
func doJob() {

	// Jobを取得
	jobMap := getAllCacheJobMap()

	// job log
	if len(jobMap) > 0 {
		log.Println("[JOB RUNNER] execute job cnt:", len(jobMap))
	}

	// jobはそれぞれ非同期で実行される
	wg := sync.WaitGroup{}
	for key, job := range jobMap {
		wg.Add(1)
		go func(key string, job *Job) {

			// panicした場合、Jobのgorutine自体が止まってしまうため
			defer catchPanic()

			defer wg.Done()

			// context timeoutを設定
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()

			// jobを実行
			defer deleteCacheJob(key)
			job.ExecFunc(ctx, job.Data)
		}(key, job)
	}

	// 全てのJobが終わるまで待機
	wg.Wait()
}

// catchPanic パニックが発生した時に、jobが強制終了されないようにする
func catchPanic() {

	if err := recover(); err != nil {

		// error log output
		switch e := err.(type) {
		case *almaerror.SystemError:
			if len(e.Params) == 0 {
				log.Println(e.Err, fmt.Sprintf("[JOB RUNNER ERROR] [SYSTEM ERROR] statuscode=%d", e.StatusCode))
			} else {
				log.Println(e.Err, fmt.Sprintf("[JOB RUNNER ERROR] statuscode=%d %v", e.StatusCode, e.Params))
			}
		case *almaerror.LogicError:
			msg := errmsg.Get("ja", e.MessageCode, e.Params...)

			// 本番は、ロジックエラー（検査例外）をログ出力しない
			if config.ConfigData.Mode != "production" {
				log.Printf("[JOB RUNNER ERROR] [LOGIC ERROR] statuscode=%d messagecode=%s message=%s",
					e.StatusCode,
					e.MessageCode,
					msg,
				)
			}

		case *almaerror.BillingError:

			log.Println(e.Err, fmt.Sprintf("[JOB RUNNER ERROR] [BILLING ERROR] messagecode=%s", e.MessageCode))
		case error:
			log.Println(e, fmt.Sprintf("[JOB RUNNER ERROR] unknown error statuscode=%d", http.StatusInternalServerError))
		default:
			// 基本ここには到達しない想定
			log.Printf("%v", err)
		}

	}

}
