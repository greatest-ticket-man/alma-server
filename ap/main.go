package main

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/mongodb/index"
	"alma-server/ap/src/infrastructure/server"
	"context"
	"flag"
)

func main() {

	// config setup
	configPath := flag.String("f", "./config/local.toml", "config file path")
	flag.Parse()
	config := config.Setup(*configPath)

	// mongodb setup
	mongodb.Setup(config.MongoDatabases)

	// mongo index setup
	index.CreateIndex(context.Background())

	// server setup
	server.Setup(config.HTTPServer)

	// run
	server.Run()
}
