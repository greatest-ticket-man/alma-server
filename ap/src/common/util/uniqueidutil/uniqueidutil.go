package uniqueidutil

import "github.com/rs/xid"

// GenerateUniqueID 一意なIDを生成 分散システムでも大丈夫
// https://qiita.com/mura-s/items/8914f34696a3502d82da
func GenerateUniqueID() string {
	guid := xid.New()
	return guid.String()
}
