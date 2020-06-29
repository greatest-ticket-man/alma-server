package main

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/infrastructure/server"
	"flag"
)

func main() {

	// config setup
	configPath := flag.String("f", "./config/local.toml", "config file path")
	flag.Parse()
	config := config.Setup(*configPath)

	// server setup
	server.Setup(config)

	// run
	server.Run()
}
