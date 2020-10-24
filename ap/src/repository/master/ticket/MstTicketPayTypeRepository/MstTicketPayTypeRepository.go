package MstTicketPayTypeRepository

import (
	"alma-server/ap/src/common/util/cacheutil"
	"alma-server/ap/src/infrastructure/mastercache"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// ThisCollectionName .
const ThisCollectionName = "MST_TICKET_PAY_TYPE"

var reflectType = reflect.TypeOf((*MstTicketPayType)(nil))

// MstTicketPayType .
type MstTicketPayType struct {
	ID         string    `bson:"_id" json:"_id"`
	Name       string    `bson:"name" json:"name"`
	Desc       string    `bson:"desc" json:"desc"`
	CreateTime time.Time `bson:"ct" json:"ct"`
	UpdateTime time.Time `bson:"ut" json:"ut"`
	// TODO 当日に支払いが必要かどうかのフラグもここに追加するようにする？
}

// LoadCache キャッシュをロードする
func LoadCache(cacheAll *gocache.Cache, dir string) {
	list := cacheutil.LoadJSON(dir, ThisCollectionName, reflectType.Elem())
	m := map[string]*MstTicketPayType{}
	for _, v := range list {
		s := v.(*MstTicketPayType)
		m[s.ID] = s
	}

	cacheutil.Set(cacheAll, ThisCollectionName, m)
}

func get() map[string]*MstTicketPayType {
	c, _ := mastercache.Cache.Get(ThisCollectionName)
	return c.(map[string]*MstTicketPayType)
}

// Get .
func Get(ticketPayTypeID string) *MstTicketPayType {
	return get()[ticketPayTypeID]
}

// GetMap .
func GetMap() map[string]*MstTicketPayType {
	return get()
}

// GetList .
func GetList() []*MstTicketPayType {

	var list []*MstTicketPayType

	mstMap := GetMap()
	for _, mst := range mstMap {
		list = append(list, mst)
	}

	return list
}
