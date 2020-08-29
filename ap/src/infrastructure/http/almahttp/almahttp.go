package almahttp

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/controller/event"
	"alma-server/ap/src/controller/hello"
	"alma-server/ap/src/controller/home/dashboard"
	"alma-server/ap/src/controller/login"
	"alma-server/ap/src/controller/member"
	"alma-server/ap/src/controller/signup"
	"alma-server/ap/src/controller/test"
	"alma-server/ap/src/controller/todo"
	"alma-server/ap/src/controller/top"
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
	n.Use(negroni.HandlerFunc(middleware.ErrorHandlingMiddleware))

	// log and recovery
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())

	// router
	router := mux.NewRouter()

	router.HandleFunc("/", top.PageHTML).Methods("GET")
	router.HandleFunc("/login", login.PageHTML).Methods("GET")
	router.HandleFunc("/login", login.Login).Methods("POST")
	router.HandleFunc("/signup", signup.PageHTML).Methods("GET")
	router.HandleFunc("/signup", signup.Signup).Methods("POST")
	router.HandleFunc("/logout", login.Logout)

	// static Staticコンテンツ
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("asset/static/")))).Methods("GET")

	// auth ログイン中のコンテンツはここ
	authRouter := mux.NewRouter().PathPrefix("/").Subrouter().StrictSlash(true)
	authRouter.HandleFunc("/home/dashboard", dashboard.PageHTML).Methods("GET")
	authRouter.HandleFunc("/home/dashboard/empty", dashboard.PageHTMLEmpty).Methods("GET")
	authRouter.HandleFunc("/event", event.PageHTML).Methods("GET")
	authRouter.HandleFunc("/event/create", event.CreatePageHTML).Methods("GET")
	authRouter.HandleFunc("/event/create", event.CreateEvent).Methods("POST")
	authRouter.HandleFunc("/event/update", event.UpdatePageHTML).Methods("GET")
	authRouter.HandleFunc("/event/update", event.UpdateEvent).Methods("POST")
	authRouter.HandleFunc("/event/list", event.GetEventList).Methods("GET")
	authRouter.HandleFunc("/member", member.PageHTML).Methods("GET")
	authRouter.HandleFunc("/hello", hello.HTML).Methods("GET")
	authRouter.HandleFunc("/test", test.PageHTML).Methods("GET")
	authRouter.HandleFunc("/todo", todo.PageHTML).Methods("GET")
	authRouter.HandleFunc("/todo/create", todo.CreateTodo).Methods("POST")
	authRouter.HandleFunc("/todo/remove", todo.RemoveTodo).Methods("POST")

	router.PathPrefix("/").Handler(negroni.New(
		negroni.HandlerFunc(middleware.AuthMiddleware),
		negroni.Wrap(authRouter),
	))

	// regist
	n.UseHandler(router)

	return n
}
