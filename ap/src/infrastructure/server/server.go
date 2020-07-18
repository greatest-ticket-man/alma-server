package server

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/infrastructure/http/almahttp"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// HTTPAPIServer api用のServer
var HTTPAPIServer *http.Server

// HTTPHTMLServer html用のServer
var HTTPHTMLServer *http.Server

// Setup ServerSetup
func Setup(config *config.HTTPServer) {

	// http api server
	HTTPAPIServer = almahttp.Setup(config)
	log.Println("api server : ", config.Address)

}

// Serve serve
func Serve(config *config.HTTPServer) {
	log.Println("http api server start !")
	go func() {

		var err error
		if config.TLS {
			err = HTTPAPIServer.ListenAndServeTLS(config.CertFile, config.KeyFile)
		} else {
			err = HTTPAPIServer.ListenAndServe()
		}
		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

}

// Shutdown .
func Shutdown() {

	// http api
	HTTPAPIServer.Shutdown(context.Background())

}

// Run Server Start
func Run(config *config.HTTPServer) {

	// serve
	Serve(config)

	// kill commandが来たら正常終了する
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("server shutdown start...")
	Shutdown()

	log.Println("Server exiting")
}
