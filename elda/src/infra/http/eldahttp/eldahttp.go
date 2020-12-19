package eldahttp

import (
	"alma-server/ap/src/common/config"
	"alma-server/elda/src/ctrl/helloctrl"
	"alma-server/elda/src/ctrl/reservectrl"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Setup .
func Setup(config *config.HTTPServer) *http.Server {

	router := Router()
	s := &http.Server{
		Handler: router,
		Addr:    config.Address,
	}
	return s
}

// Router .
func Router() *negroni.Negroni {
	n := negroni.New()

	// n.Use(negroni.HandlerFunc(middleware.ErrorHandlingMiddleware))

	// log and recovery
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery()) // TODO 自分でrecoveryをmiddlewareで実装する

	// router
	router := mux.NewRouter()

	router.HandleFunc("/", helloctrl.PageHTML).Methods("GET")
	router.HandleFunc("/reserve", reservectrl.PageHTML).Methods("GET")

	// regist
	n.UseHandler(router)
	return n
}
