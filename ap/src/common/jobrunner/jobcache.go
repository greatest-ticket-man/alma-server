package jobrunner

import (
	"time"

	"github.com/patrickmn/go-cache"
)

var cacheJob *cache.Cache

// cacheSetUp job用のキャッシュのセットアップ
func cacheSetUp() {
	// 初期化
	cacheJob = cache.New(10*time.Minute, 10*time.Minute)
}

// getAllCacheJobMap キャッシュをMapで取得する
func getAllCacheJobMap() map[string]*Job {

	cacheMap := cacheJob.Items()
	resultMap := map[string]*Job{}
	for key, cache := range cacheMap {
		resultMap[key] = cache.Object.(*Job)
	}

	return resultMap
}

// setCacheJob Jobキャッシュにデータを追加する
func setCacheJob(k string, v interface{}) {
	cacheJob.Set(k, v, cache.DefaultExpiration)
}

// deleteCacheJob Jobキャッシュのデータを削除する
func deleteCacheJob(k string) {
	cacheJob.Delete(k)
}
