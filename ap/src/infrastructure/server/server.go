package server

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/common/error/chk"
	"alma-server/ap/src/infrastructure/http/almahttp"
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
)

// HTTPServer Server
var HTTPServer *http.Server

// Setup ServerSetup
func Setup(config *config.HTTPServer) {

	// http api server
	HTTPServer = almahttp.Setup(config)
	log.Println("api server : ", config.Address)

}

// Serve serve
func Serve(config *config.HTTPServer) {
	log.Println("http api server start !")
	go func() {

		var err error
		if config.TLS {
			err = HTTPServer.ListenAndServeTLS(config.CertFile, config.KeyFile)
		} else {
			err = HTTPServer.ListenAndServe()
		}
		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

}

// TestServe テスト時は、ポートがバッティングするため
// テスト時のみ、Unix Domain Socket で立ち上げる
func TestServe() {

	// httpUnixDomainSocketPath = "/tmp/alma-ap-test-http-udx-"
	httpUnixDomainSocketPath := fmt.Sprintf("/tmp/alma-ap-test-http-udx-%s.sock", uuid.New().String())

	log.Println("Http Server Start :", httpUnixDomainSocketPath)

	go func() {

		tcpServer, err := net.Listen("unix", httpUnixDomainSocketPath)
		chk.SE(err)

		err = HTTPServer.Serve(tcpServer)
		chk.SE(err)
	}()

}

// Shutdown .
func Shutdown() {

	// http api
	HTTPServer.Shutdown(context.Background())
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
