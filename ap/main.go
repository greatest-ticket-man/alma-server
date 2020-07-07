package main

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/mongodb/index"
	"alma-server/ap/src/infrastructure/server"
	"alma-server/ap/src/infrastructure/stripe"
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

	// stripe
	stripe.Setup(config.Stripe)

	// server setup
	server.Setup(config.HTTPServer)

	// run
	server.Run()
}
