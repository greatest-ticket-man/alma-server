package MstMenuRepository

import (
	"alma-server/ap/src/common/util/cacheutil"
	"alma-server/ap/src/infrastructure/mastercache"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// ThisCollectionName .
const ThisCollectionName = "MST_MENU"

var reflectType = reflect.TypeOf((*MstMenu)(nil))

// MstMenu SideMenuなどで表示するデータ
type MstMenu struct {
	ID         string    `bson:"_id" json:"_id"`           // 識別ID unique
	Title      string    `bson:"title" json:"title"`       // 表示する名前
	Desc       string    `bson:"desc" json:"desc"`         // 説明
	Parent     string    `bson:"parent" json:"parent"`     // 親のID
	Children   []string  `bson:"children" json:"children"` // 子どもたちのID
	Path       string    `bson:"path" json:"path"`         // クリックしたときに遷移するpath
	Icon       string    `bson:"icon" json:"icon"`         // Icon
	CreateTime time.Time `bson:"ct" json:"ct"`
	UpdateTime time.Time `bson:"ut" json:"ut"`
}

// LoadCache キャッシュをロードする
func LoadCache(cacheAll *gocache.Cache, dir string) {

	list := cacheutil.LoadJSON(dir, ThisCollectionName, reflectType.Elem())
	m := map[string]*MstMenu{}
	for _, v := range list {
		s := v.(*MstMenu)
		m[s.ID] = s
	}
	cacheutil.Set(cacheAll, ThisCollectionName, m)
}

func get() map[string]*MstMenu {
	c, _ := mastercache.Cache.Get(ThisCollectionName)
	return c.(map[string]*MstMenu)
}

// GetMap .
func GetMap() map[string]*MstMenu {
	return get()
}
