package test

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/projectpathap"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/server"
	"alma-server/ap/src/infrastructure/stripe"
	"fmt"
)

// Setup testのためのセットアップ
func Setup() {

	// localpath
	path := fmt.Sprintf("%s/config/local.toml", projectpathap.GetRoot())

	// config
	config := config.Setup(path)

	// mongodb
	mongodb.Setup(config.MongoDatabases)

	// stripe
	stripe.Setup(config.Stripe)

	// server
	server.Setup(config.HTTPServer)

	// start
	server.Serve(config.HTTPServer)

}
