package MstEventRoleRepository

import (
	"alma-server/ap/src/common/util/cacheutil"
	"alma-server/ap/src/infrastructure/mastercache"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// reflectType .
var reflectType = reflect.TypeOf((*MstEventRole)(nil))

// ThisCollectionName .
const ThisCollectionName = "MST_EVENT_ROLE"

// MstEventRole どんな役割があるか？
type MstEventRole struct {
	ID         string    `bson:"_id,omitempty" json:"_id"`
	Name       string    `bson:"name" json:"name"`
	CreateTime time.Time `bson:"ct" json:"ct"`
	UpdateTime time.Time `bson:"ut" json:"ut"`
	// TODO どのように役割を制御するか？ pathごとに権限をつける感じが一番いいような気がする
}

// LoadCache キャッシュをロードする
func LoadCache(cacheAll *gocache.Cache, dir string) {
	list := cacheutil.LoadJSON(dir, ThisCollectionName, reflectType.Elem())
	m := map[string]*MstEventRole{}
	for _, v := range list {
		s := v.(*MstEventRole)
		m[s.ID] = s
	}
	cacheutil.Set(cacheAll, ThisCollectionName, m)
}

// Get .
func Get(roleID string) *MstEventRole {
	c, _ := mastercache.Cache.Get(ThisCollectionName)
	return c.(map[string]*MstEventRole)[roleID]
}
