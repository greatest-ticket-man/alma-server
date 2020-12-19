package main

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/logger"
	"alma-server/ap/src/infrastructure/mongodb"
	"alma-server/ap/src/infrastructure/mongodb/index"
	"alma-server/elda/src/infra/server"
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
	conf := config.Setup(*configPath)

	// mongo setup
	mongodb.Setup(conf.MongoDatabases)

	// mongo index setup
	index.CreateIndex(context.Background())

	// TODO server
	// eldahttp.Setup(conf.HTTPServer)
	server.Setup(conf.HTTPServer)

	server.Run(conf.HTTPServer)

}
