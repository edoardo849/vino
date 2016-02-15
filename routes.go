package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GameIndex",
		"GET",
		"/games",
		GameIndex,
	},
	Route{
		"GameCreate",
		"POST",
		"/games",
		GameCreate,
	},
	Route{
		"GameShow",
		"GET",
		"/games/{id}",
		GameShow,
	},
	Route{
		"GameGuessLetter",
		"POST",
		"/games/{id}",
		GameGuessLetter,
	},
}
