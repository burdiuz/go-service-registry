package router

import (
	"net/http"

	matcher "../../path/matcher"
)

func undefinedPath(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))
}

// Router object contains paths and handler functions for endpoints
type Router struct {
	paths     *matcher.PathRegistry
	undefined matcher.PathHandler
}

// New creates new router instance with its own paths mapping
func New(undefined matcher.PathHandler) *Router {
	router := Router{paths: matcher.PathRegistryNew(), undefined: undefined}

	return &router
}

// Route Add route handler
func (r *Router) Route(path string, getHandler matcher.PathHandler) *Route {
	route := RouteNew(getHandler, r.undefined)

	r.paths.Add(path, route.GetHandler())

	if getHandler != nil {
		route.Get(getHandler)
	}

	return route
}

// GetHandler returns function that handles HTTP requests using router
func (r *Router) GetHandler() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		match := r.paths.Get(request.URL.Path)

		switch {
		case match != nil:
			match.Handler(writer, request, match.Vars)
		case r.undefined != nil:
			r.undefined(writer, request, nil)
		case match != nil:
			undefinedPath(writer, request)
		}
	}
}
