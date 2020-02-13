package router

import (
	"fmt"
	"net/http"

	matcher "../go-url-path-matcher"
)

// Router object contains paths and handler functions for endpoints
type Router struct {
	paths     *matcher.PathRegistry
	undefined matcher.PathHandler
}

// New creates new router instance with its own paths mapping
func New(notFound matcher.PathHandler) *Router {
	router := Router{paths: matcher.NewPathRegistry(), undefined: notFound}

	return &router
}

/*
  TODO Add ParamsRoute struct which will work with ParamsPath,
  Router checks if path contains /: and uses Params*, if not -- uses normal Route and Path.

*/

// Route Add route handler
func (r *Router) Route(path string, getHandler matcher.PathHandler) *Route {
	route := NewRoute(getHandler, r.undefined)

	r.paths.Add(path, route.GetHandler())

	return route
}

// GetHandler returns function that handles HTTP requests using router
func (r *Router) GetHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		match := r.paths.Get(request.URL.Path)

		fmt.Printf("%v, %v\n", request.URL.Path, match)

		switch {
		case match != nil:
			match.Handler(writer, request, match.Params)
		case r.undefined != nil:
			r.undefined(writer, request, nil)
		case match != nil:
			http.NotFound(writer, request)
		}
	}
}
