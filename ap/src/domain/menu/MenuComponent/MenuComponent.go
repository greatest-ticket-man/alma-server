package MenuComponent

import (
	"alma-server/ap/src/infrastructure/grpc/proto/menu"
	"alma-server/ap/src/repository/master/menu/MstMenuRepository"
)

// CreateMenuInfo .
func CreateMenuInfo(mstMenu *MstMenuRepository.MstMenu, selectedMenuID string, childMenuInfoList []*menu.MenuInfo) *menu.MenuInfo {

	return &menu.MenuInfo{
		Title:             mstMenu.Title,
		Desc:              mstMenu.Desc,
		Path:              mstMenu.Path,
		Icon:              mstMenu.Icon,
		IsSelected:        mstMenu.ID == selectedMenuID,
		ChildMenuInfoList: childMenuInfoList,
	}

}
