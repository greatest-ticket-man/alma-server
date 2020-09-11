package almafile

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/common/projectpathap"
	"net/http"
	"path/filepath"

	// static
	_ "alma-server/ap/statik/static"

	// template
	_ "alma-server/ap/statik/template"

	// asset
	_ "alma-server/ap/statik"

	"github.com/rakyll/statik/fs"
)

// GetStaticFileSystem .
func GetStaticFileSystem(config *config.AlmaConfig) http.FileSystem {

	// localだけ、fileを読み込み
	if config.Mode == "local" {
		return http.Dir(filepath.Join(projectpathap.GetRoot(), "asset", "static"))
	}

	// one binaryを読み込み
	s, err := fs.NewWithNamespace("static")
	chk.SE(err)
	return s
}

// GetAssetFileSystem .
func GetAssetFileSystem(config *config.AlmaConfig) http.FileSystem {

	// localだけ、fileを読み込む
	if config.Mode == "local" {
		return http.Dir(filepath.Join(projectpathap.GetRoot(), "asset"))
	}

	// one binary を読み込む
	s, err := fs.NewWithNamespace("asset")
	chk.SE(err)
	return s
}
