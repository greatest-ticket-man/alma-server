package AccountComponent

import "alma-server/ap/src/repository/user/UserAccountRepository"

// GetUserAccountMap ユーザーのアカウント情報をMapにする
func GetUserAccountMap(list []*UserAccountRepository.UserAccount) map[string]*UserAccountRepository.UserAccount {
	m := map[string]*UserAccountRepository.UserAccount{}
	for _, v := range list {
		m[v.ID] = v
	}
	return m
}
