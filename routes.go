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
		"Ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		"QueueAdd",
		"POST",
		"/queue/add",
		QueueAdd,
	},
	Route{
		"QueueNext",
		"POST",
		"/queue/next",
		QueueNext,
	},
	Route{
		"QueueClear",
		"POST",
		"/queue/clear",
		QueueClear,
	},
	Route{
		"QueueTop",
		"GET",
		"/queue/top",
		QueueGetTop,
	},
	Route{
		"QueueGet",
		"GET",
		"/queue",
		QueueGet,
	},
	Route{
		"StreamGetId",
		"GET",
		"/stream/{id}",
		GetStreamId,
	},
	Route{
		"StreamGet",
		"GET",
		"/stream",
		GetStream,
	},
	Route{
		"Search",
		"POST",
		"/search",
		Search,
	},
}
