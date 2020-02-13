package router

import (
	"errors"
	"fmt"
	"net/http"

	matcher "../go-url-path-matcher"
	utils "./utils"
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

/*NewRoute creates new Route and accepts two parameters
- get is a function to handle GET requests
- undefined is a function to handle requests to methods that were left with no handlers
*/
func NewRoute(get interface{}, undefined interface{}) *Route {
	route := Route{}

	if undefined != nil {
		route.undefined = utils.HandleCustom(undefined)
	}

	if get != nil {
		route.Get(get)
	}

	return &route
}

// AddMethod adds handler for specified HTTP method
func (r *Route) AddMethod(method string, handler matcher.PathHandler) error {
	if method == "" || handler == nil {
		return errors.New("Method and Handler are required parameters.")
	}

	if r.methods == nil {
		r.methods = make(map[string]matcher.PathHandler)
	}

	if r.methods[method] != nil {
		return fmt.Errorf("Handler for method %q is already registered", method)
	}

	r.methods[method] = handler

	return nil
}

/*AddCustom allows to use functions of custom signature for routing, these signatures supported
func()
func(writer http.ResponseWriter)
func(writer http.ResponseWriter, request *http.Request)
func(writer http.ResponseWriter, request *http.Request, params matcher.PathParams)
Could be useful to reuse handlers that were applied to http.HandleFunc() or for paths
with no parameters just ignore receiving PathParams.
*/
func (r *Route) AddCustom(method string, handler interface{}) error {
	custom := utils.HandleCustom(handler)
	return r.AddMethod(method, custom)
}

func (r *Route) addOrPanic(method string, handler interface{}) *Route {
	err := r.AddCustom(method, handler)

	if err != nil {
		panic(err)
	}

	return r
} // Get adds handler for GET HTTP method
func (r *Route) Get(handler interface{}) *Route {
	return r.addOrPanic("GET", handler)
}

// Head adds handler for HEAD HTTP method
func (r *Route) Head(handler interface{}) *Route {
	return r.addOrPanic("HEAD", handler)
}

// Post adds handler for POST HTTP method
func (r *Route) Post(handler interface{}) *Route {
	return r.addOrPanic("POST", handler)
}

// Put adds handler for PUT HTTP method
func (r *Route) Put(handler interface{}) *Route {
	return r.addOrPanic("PUT", handler)
}

// Patch adds handler for PATCH HTTP method
func (r *Route) Patch(handler interface{}) *Route {
	return r.addOrPanic("PATCH", handler)
}

// Delete adds handler for DELETE HTTP method
func (r *Route) Delete(handler interface{}) *Route {
	return r.addOrPanic("DELETE", handler)
}

// Delete adds handler for DELETE HTTP method
func (r *Route) Options(handler interface{}) *Route {
	return r.addOrPanic("OPTIONS", handler)
}

// HasMethod checks if HTTP method has handler registered
func (r *Route) HasMethod(method string) bool {
	return r.methods[method] != nil
}

// Call calls handler depending on HTTP method from http.Request.Method or handler for unknown method
func (r *Route) Call(writer http.ResponseWriter, request *http.Request, params matcher.PathParams) {
	switch {
	case r.methods[request.Method] != nil:
		r.methods[request.Method](writer, request, params)
	case r.undefined != nil:
		r.undefined(writer, request, params)
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
	return func(writer http.ResponseWriter, request *http.Request, params matcher.PathParams) {
		r.Call(writer, request, params)
	}
}
