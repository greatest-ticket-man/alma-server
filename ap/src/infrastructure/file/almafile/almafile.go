package almafile

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"net/http"

	// static file
	_ "alma-server/ap/statik/static"

	// template file
	_ "alma-server/ap/statik/template"

	"github.com/rakyll/statik/fs"
)

// GetStaticFileSystem .
func GetStaticFileSystem(config *config.AlmaConfig) http.FileSystem {

	// localだけ、fileを読み込み
	if config.Mode == "local" {
		return http.Dir("asset/static/")
	}

	// one binaryを読み込み
	s, err := fs.NewWithNamespace("asset/static/")
	chk.SE(err)
	return s
}

// GetTemplateFileSystem .
func GetTemplateFileSystem(config *config.AlmaConfig) http.FileSystem {

	// localだけ、fileを読み込む
	if config.Mode == "local" {
		return http.Dir("asset/template/")
	}

	// one binaryを読み込み
	s, err := fs.NewWithNamespace("asset/template/")
	chk.SE(err)
	return s
}
