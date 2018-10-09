package main

import (
	"github.com/gorilla/mux"
	"net/http"

	"log"
	"time"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(_handler(route.HandlerFunc, route.Name))
	}

	return router
}

func _handler(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_logRequest(r, name)
		inner.ServeHTTP(w, r)
	})
}

func _logRequest(r *http.Request, routeName string) {
	start := time.Now()

	log.Printf(
		"%s\t%s\t%s\t%s\t",
		r.Method,
		r.RequestURI,
		routeName,
		time.Since(start),
	)
}
