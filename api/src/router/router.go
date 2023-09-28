package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

// Togenerate will return a router with the configured routes
func Togenerate() *mux.Router {
	r := mux.NewRouter()
	routes.Setup(r)
	return r
}
