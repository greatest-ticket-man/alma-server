package mastercache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// Cache .
var Cache *gocache.Cache

// キャッシュの期限は設定しない、クリーンアップは5分ごと
func init() {
	Cache = gocache.New(gocache.NoExpiration, 5*time.Minute)
}
