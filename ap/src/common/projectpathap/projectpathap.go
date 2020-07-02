package projectpathap

import (
	"path/filepath"
	"runtime"
)

// Testの実行時などに、Rootディレクトリの取得が簡単にはできない
// だが、下記のハック的な書き方をすると、main実行でもtest実行でも
// Rootディレクトリが取得できる。

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../../../")
)
