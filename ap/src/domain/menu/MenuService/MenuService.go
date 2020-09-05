package MenuService

import (
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/domain/menu/MenuComponent"
	"alma-server/ap/src/infrastructure/grpc/proto/menu"
	"alma-server/ap/src/repository/master/menu/MstMenuRepository"
	"errors"
	"time"
)

// GetMenu メニューを取得する
// 無限ループが発生する可能性があるので、取得に3秒以上かかった場合はPanicする
func GetMenu(menuID string) *menu.MenuInfo {
	mstMenuMap := MstMenuRepository.GetMap()

	ch := make(chan *menu.MenuInfo)
	go func() {
		ch <- getMenu(menuID, mstMenuMap)
	}()

	select {
	case result := <-ch:
		return result
	case <-time.After(time.Duration(3) * time.Second):
		chk.SE(errors.New("Menuの取得時にタイムアウトしました"))
		return nil
	}
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
