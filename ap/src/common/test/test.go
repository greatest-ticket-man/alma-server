package test

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/projectpathap"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/server"
	"fmt"
)

// Setup testのためのセットアップ
func Setup() {

	// localpath
	path := fmt.Sprintf("%s/config/local.toml", projectpathap.Root)

	// config
	config := config.Setup(path)

	// mongodb
	mongodb.Setup(config.MongoDatabases)

	// server
	server.Setup(config.HTTPServer)

	// start
	server.Serve()

}
