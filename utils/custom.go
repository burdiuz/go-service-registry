package utils

import (
	"net/http"
	"reflect"

	matcher "../path/matcher"
)

/*HandleCustom allows using functions with these signatures:
func()
func(res http.ResponseWriter)
func(res http.ResponseWriter, req *http.Request)
func(res http.ResponseWriter, req *http.Request, params matcher.PathParams)
*/
func HandleCustom(customHandler interface{}) matcher.PathHandler {
	fnType := reflect.TypeOf(customHandler)

	if fnType.Kind() != reflect.Func {
		panic("Custom handler can be only of type Function.")
	}

	switch fnType.NumIn() {
	case 0:
		fn := customHandler.(func())
		return func(_ http.ResponseWriter, _ *http.Request, _ matcher.PathParams) {
			fn()
		}
	case 1:
		fn := customHandler.(func(res http.ResponseWriter))
		return func(res http.ResponseWriter, _ *http.Request, _ matcher.PathParams) {
			fn(res)
		}
	case 2:
		fn := customHandler.(func(res http.ResponseWriter, req *http.Request))
		return func(res http.ResponseWriter, req *http.Request, _ matcher.PathParams) {
			fn(res, req)
		}
	case 3:
		fn := customHandler.(func(res http.ResponseWriter, req *http.Request, params matcher.PathParams))
		return func(res http.ResponseWriter, req *http.Request, params matcher.PathParams) {
			fn(res, req, params)
		}
	default:
		panic("Custom handler got an unexpected count of input parameters.")
	}
}
