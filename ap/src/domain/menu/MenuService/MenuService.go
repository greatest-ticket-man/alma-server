package MenuService

import (
	"alma-server/ap/src/domain/menu/MenuComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/menu"
	"alma-server/ap/src/repository/master/menu/MstMenuRepository"
)

// GetMenu .
func GetMenu(menuID string) *menu.MenuInfo {
	mstMenuMap := MstMenuRepository.GetMap()

	// TODO 無限ループを感知するためにchanを使う

	return getMenu(menuID, mstMenuMap)
}

func getMenu(menuID string, mstMenuMap map[string]*MstMenuRepository.MstMenu) *menu.MenuInfo {

	mstMenu := mstMenuMap[menuID]
	if mstMenu == nil {
		return nil
	}

	var childMenuInfoList []*menu.MenuInfo
	for _, childMenuID := range mstMenu.Children {
		childMenuInfoList = append(childMenuInfoList, getMenu(childMenuID, mstMenuMap))
	}

	return MenuComponent.CreateMenuInfo(mstMenu, childMenuInfoList)
}
