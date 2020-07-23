package almahttp

import (
	"alma-server/ap/src/common/config"
	"alma-server/ap/src/controller/dashboard"
	"alma-server/ap/src/controller/hello"
	"alma-server/ap/src/controller/login"
	"alma-server/ap/src/controller/todo"
	"alma-server/ap/src/controller/top"
	"alma-server/ap/src/infrastructure/http/middleware"
	"net/http"
	"path/filepath"

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
	// n.Use(negroni.NewRecovery())

	// router
	router := mux.NewRouter()

	router.HandleFunc("/", top.PageHTML).Methods("GET")
	router.HandleFunc("/dashboard", dashboard.PageHTML).Methods("GET")

	router.HandleFunc("/api/hello", hello.API).Methods("GET")

	router.HandleFunc("/hello", hello.HTML).Methods("GET")

	router.HandleFunc("/login", login.PageHTML).Methods("GET")

	router.HandleFunc("/todo", todo.PageHTML).Methods("GET")
	router.HandleFunc("/todo/create", todo.CreateTodo).Methods("POST")
	router.HandleFunc("/todo/remove", todo.RemoveTodo).Methods("POST")

	// static
	router.PathPrefix("/static").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("asset/static/")))).Methods("GET")

	// regist
	n.UseHandler(router)

	return n
}

// neuteredFileSystem file server
type neuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {

	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}

			return nil, err
		}
	}

	return f, nil
}
