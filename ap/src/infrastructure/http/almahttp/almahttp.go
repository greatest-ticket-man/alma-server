package almahttp

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/controller/hello"
	"alma-server/ap/src/controller/login"
	"alma-server/ap/src/controller/todo"
	"alma-server/ap/src/infrastructure/http/middleware"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

// Setup setup
func Setup(config *config.HTTPServer) *http.Server {

	router := Router()
	s := &http.Server{
		Handler: router,
		Addr:    config.Address,
	}

	return s
}

// Router router handling
func Router() *negroni.Negroni {

	n := negroni.New()

	// middleware
	n.Use(negroni.HandlerFunc(middleware.CorsMiddleware))

	// TODO error catch wrapper

	// log and recovery
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())

	// router
	router := mux.NewRouter()

	// hello
	// router.HandleFunc("/api/hello", HelloController.Hello).Methods(methodsGet)

	// // html
	// router.HandleFunc("/hello", HelloController.HelloHTML).Methods(methodsGet)

	router.HandleFunc("/api/hello", hello.API).Methods("GET")

	router.HandleFunc("/hello", hello.HTML).Methods("GET")

	router.HandleFunc("/login", login.PageHTML).Methods("GET")

	router.HandleFunc("/todo", todo.PageHTML).Methods("GET")

	// regist
	n.UseHandler(router)

	return n
}
