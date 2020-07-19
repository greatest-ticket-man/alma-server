package projectpathap

import (
	"alma-server/ap/src/common/config"
	"path/filepath"
	"runtime"
)

// Testの実行時などに、Rootディレクトリの取得が簡単にはできない
// だが、下記のハック的な書き方をすると、main実行でもtest実行でも
// Rootディレクトリが取得できる。

var (
	_, b, _, _ = runtime.Caller(0)

	// TestRoot .
	TestRoot = filepath.Join(filepath.Dir(b), "../../../")
)

// GetRoot Rootディレクトリを取得する
func GetRoot() string {

	rootDirectory := config.ConfigData.RootDirectory

	if rootDirectory == "" {
		return filepath.Join(filepath.Dir(b), "../../../")
	}

	return rootDirectory
}
