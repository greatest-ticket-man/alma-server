package cacheutil

import (
	"alma-server/ap/src/common/error/chk"
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"time"

	gocache "github.com/patrickmn/go-cache"
)

// New .
func New() *gocache.Cache {
	return gocache.New(gocache.NoExpiration, 5*time.Minute)
}

// Set .
func Set(cacheAll *gocache.Cache, key string, v interface{}) {
	cacheAll.Set(key, v, gocache.NoExpiration)
}

// LoadJSON .
// reflectTypeはポインタのreflectTypeではだめ
func LoadJSON(dir string, colName string, reflectType reflect.Type) []interface{} {
	return Load(filepath.Join(dir, colName+".json"), reflectType)
}

// Load ファイルから取得する
func Load(p string, reflectType reflect.Type) []interface{} {

	f, err := os.Open(p)
	chk.SE(err)
	defer f.Close()

	var list []interface{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		str := scanner.Text()
		if str == "" {
			continue
		}

		v := reflect.New(reflectType)
		json.Unmarshal([]byte(str), v.Interface())
		list = append(list, v.Interface())
	}

	chk.SE(scanner.Err())
	return list
}
