package main

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/jobrunner"
	"alma-server/ap/src/common/jwt"
	"alma-server/ap/src/common/logger"
	"alma-server/ap/src/infrastructure/mastercache/cacheall"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/mongodb/index"
	"alma-server/ap/src/infrastructure/prometheus"
	"alma-server/ap/src/infrastructure/server"
	"alma-server/ap/src/infrastructure/stripe"
	"context"
	"flag"
)

var (
	hash      string
	builddate string
	goversion string
	goos      string
	goarch    string
)

func main() {

	flag.Parse()
	logger.Info(goversion)
	logger.Infof("goos/goarch=%s/%s\n", goos, goarch)
	logger.Infof("githash=%s\n", hash)
	logger.Infof("build at %s\n", builddate)

	// config setup
	configPath := flag.String("f", "./config/local.toml", "config file path")
	flag.Parse()
	config := config.Setup(*configPath)

	// jwt setup
	jwt.Setup()

	// mongodb setup
	mongodb.Setup(config.MongoDatabases)

	// mongo index setup
	index.CreateIndex(context.Background())

	// stripe
	stripe.Setup(config.Stripe)

	// cache setup
	cacheall.LoadMaster(config.MasterCacheDir)

	// jobrunner
	jobrunner.Run()

	// prometheus server setup
	prometheus.Setup()

	// server setup
	server.Setup(config.HTTPServer)

	// run
	server.Run(config.HTTPServer)
}
