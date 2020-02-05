package router

import (
	"fmt"
	"net/http"

	matcher "../../path/matcher"
)

func undefinedMethod(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(405)
	w.Write([]byte("Method Not Allowed"))
}

// Route represents one endpoint or route on server, it contains handlers for various HTTP methods and calls them on request
type Route struct {
	/*

	  TODO I have to use reflection to find out how many aruments funtion accepts to support standard HTTP handlers and PathHandler's

	*/
	methods   map[string]matcher.PathHandler
	undefined matcher.PathHandler
}

func RouteNew(get matcher.PathHandler, undefined matcher.PathHandler) *Route {
	route := Route{undefined: undefined}

	if get != nil {
		route.Get(get)
	}

	return &route
}

// Get adds handler for GET HTTP method
func (r *Route) Get(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("GET", handler)
}

// Head adds handler for HEAD HTTP method
func (r *Route) Head(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("HEAD", handler)
}

// Post adds handler for POST HTTP method
func (r *Route) Post(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("POST", handler)
}

// Put adds handler for PUT HTTP method
func (r *Route) Put(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("PUT", handler)
}

// Patch adds handler for PATCH HTTP method
func (r *Route) Patch(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("PATCH", handler)
}

// Delete adds handler for DELETE HTTP method
func (r *Route) Delete(handler matcher.PathHandler) (*Route, error) {
	return r.AddMethod("DELETE", handler)
}

// AddMethod adds handler for specified HTTP method
func (r *Route) AddMethod(method string, handler matcher.PathHandler) (*Route, error) {
	if r.methods == nil {
		r.methods = make(map[string]matcher.PathHandler)
	}

	if r.methods[method] != nil {
		return nil, fmt.Errorf("Handler for method %q is already registered", method)
	}

	r.methods[method] = handler

	return r, nil
}

// HasMethod checks if HTTP method has handler registered
func (r *Route) HasMethod(method string) bool {
	return r.methods[method] != nil
}

// Call calls handler depending on HTTP method from http.Request.Method or handler for unknown method
func (r *Route) Call(writer http.ResponseWriter, request *http.Request, vars matcher.PathVars) {
	switch {
	case r.methods[request.Method] != nil:
		r.methods[request.Method](writer, request, vars)
	case r.undefined != nil:
		r.undefined(writer, request, vars)
	default:
		undefinedMethod(writer, request)
	}
}

// SetUndefined saves handler that will be called if n ohandler registered for HTTP method
func (r *Route) SetUndefined(handler matcher.PathHandler) {
	r.undefined = handler
}

// GetHandler return HTTP response handler that will pick handlers from a route depending on HTTP method
func (r *Route) GetHandler() matcher.PathHandler {
	return func(writer http.ResponseWriter, request *http.Request, vars matcher.PathVars) {
		r.Call(writer, request, vars)
	}
}
