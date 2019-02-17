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
	// util
	Route{
		"Ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		"Version",
		"GET",
		"/version",
		GetVersion,
	},
	// queue
	Route{
		"QueueAdd",
		"POST",
		"/queue/add",
		QueueAdd,
	},
	Route{
		"QueueMove",
		"POST",
		"/queue/move",
		QueueMove,
	},
	Route{
		"QueueRemove",
		"POST",
		"/queue/remove",
		QueueRemove,
	},
	Route{
		"QueueClear",
		"POST",
		"/queue/clear",
		QueueClear,
	},
	Route{
		"QueueGet",
		"GET",
		"/queue",
		QueueGet,
	},
	// stream
	Route{
		"StreamGet",
		"GET",
		"/stream/{id}",
		GetStream,
	},
	// search
	Route{
		"Search",
		"POST",
		"/search",
		Search,
	},
	Route{
		"Info",
		"GET",
		"/info/{youtubeId}",
		Info,
	},
}
