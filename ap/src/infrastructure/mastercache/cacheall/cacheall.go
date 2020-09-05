package cacheall

import (
	"alma-server/ap/src/infrastructure/mastercache"
	"alma-server/ap/src/repository/master/authority/MstEventAuthRepository"
	"alma-server/ap/src/repository/master/authority/MstEventRoleRepository"
	"alma-server/ap/src/repository/master/menu/MstMenuRepository"
)

// LoadMaster .
func LoadMaster(dir string) {
	MstEventAuthRepository.LoadCache(mastercache.Cache, dir)
	MstEventRoleRepository.LoadCache(mastercache.Cache, dir)
	MstMenuRepository.LoadCache(mastercache.Cache, dir)
}
