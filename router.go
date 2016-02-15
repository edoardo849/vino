package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*NewRouter is the mux Router from the
 * Gorilla Web Toolkit. It also takes a
 * handler for logging incoming requests
 * to the console. Production-ready routers
 * should log these into some kind of persistent storage
 */
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
