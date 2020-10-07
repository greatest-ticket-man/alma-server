package prometheus

import (
	"alma-server/ap/src/common/config"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	negronimiddleware "github.com/slok/go-http-metrics/middleware/negroni"
	"github.com/urfave/negroni"
)

// PrometheusServer Server
var PrometheusServer *http.Server

// Setup ServerSetup
func Setup() {
	PrometheusServer = &http.Server{
		Handler: promhttp.Handler(),
		Addr:    config.ConfigData.PrometheusServer.Address,
	}
	log.Println("prometheus metrics : ", config.ConfigData.PrometheusServer.Address)
}

// NegroniMiddleware negroni用 Handler
func NegroniMiddleware() negroni.Handler {

	mdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	return negronimiddleware.Handler("", mdlw)
}

// Serve prometheus metrics server
func Serve() {

	log.Println("http prometheus ap server start !")

	c := config.ConfigData.PrometheusServer

	go func() {
		var err error
		if c.TLS {
			err = PrometheusServer.ListenAndServeTLS(c.CertFile, c.KeyFile)
		} else {
			err = PrometheusServer.ListenAndServe()
		}

		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()
}
