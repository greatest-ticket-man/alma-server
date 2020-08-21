package MstEventAuthRepository

import (
	"alma-server/ap/src/common/util/cacheutil"
	"alma-server/ap/src/infrastructure/mastercache"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var reflectType = reflect.TypeOf((*MstEventAuth)(nil))

// var reflectType = reflect.TypeOf(&MstEventAuth{})

// ThisCollectionName .
const ThisCollectionName = "MST_EVENT_AUTH"

// MstEventAuth イベントで使用される権限の種類
// Role(役割)を組み合わせて、細かな権限を作成できるようにする
type MstEventAuth struct {
	ID         string          `bson:"_id,omitempty" json:"_id"`
	Name       string          `bson:"name" json:"name"`   // 権限の名前
	RoleMap    map[string]bool `bson:"roles" json:"roles"` // 役割の詳細 key: roleID, value: 有効かどうか
	IsRoot     bool            `bson:"root" json:"root"`
	CreateTime time.Time       `bson:"ct" json:"ct"`
	UpdateTime time.Time       `bson:"ut" json:"ut"`
}

// LoadCache キャッシュをロードする
func LoadCache(cacheAll *gocache.Cache, dir string) {
	list := cacheutil.LoadJSON(dir, ThisCollectionName, reflectType.Elem())
	m := map[string]*MstEventAuth{}
	for _, v := range list {
		s := v.(*MstEventAuth)
		m[s.ID] = s
	}
	cacheutil.Set(cacheAll, ThisCollectionName, m)
}

// Get .
func Get(eventAuthID string) *MstEventAuth {
	c, _ := mastercache.Cache.Get(ThisCollectionName)
	return c.(map[string]*MstEventAuth)[eventAuthID]
}
