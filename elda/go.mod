module alma-server/elda

go 1.15

replace alma-server/ap => ../ap

require (
	alma-server/ap v0.0.0-00010101000000-000000000000
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.8.0
	github.com/urfave/negroni v1.0.0
)
