package main

import (
	"net/http"
)

// non-authenticated route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{
		"QueueAdd",
		"POST",
		"/queue/add",
		QueueAdd,
	},
}
