package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route defining a Route structure
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes deifning an new Routes type
type Routes []Route

// NewRouter this function returns a mux router
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

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"TodoIndex",
		"GET",
		"/todos",
		TodoIndex,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoShow,
	},
	Route{
		"CreateAccount",
		"GET",
		"/createAccount/{name}",
		CreateAccount,
	},
	Route{
		"GetAccount",
		"GET",
		"/getAccount/{name}",
		GetAccount,
	},
}
