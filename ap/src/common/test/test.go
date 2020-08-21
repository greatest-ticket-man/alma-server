package test

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/jobrunner"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/projectpathap"
	"alma-server/ap/src/infrastructure/mastercache/cacheall"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/server"
	"alma-server/ap/src/infrastructure/stripe"
	"fmt"
	"path/filepath"
)

// Setup testのためのセットアップ
func Setup() {

	// localpath
	path := fmt.Sprintf("%s/config/local.toml", projectpathap.TestRoot)

	// config
	config := config.Setup(path)

	// jwt
	jwt.Setup()

	// mongodb
	mongodb.Setup(config.MongoDatabases)

	// stripe
	stripe.Setup(config.Stripe)

	// cache setup
	cacheall.LoadMaster(filepath.Join(projectpathap.TestRoot, config.MasterCacheDir))

	// jobrunner
	jobrunner.Run()

	// server
	server.Setup(config.HTTPServer)

	// start
	// server.Serve(config.HTTPServer)
	server.TestServe()

}
