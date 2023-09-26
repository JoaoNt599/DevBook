package router

import "github.com/gorilla/mux"

// Togenerate will return a router with the configured routes
func Togenerate() *mux.Router {
	return mux.NewRouter()
}
