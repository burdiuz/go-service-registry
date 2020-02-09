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
	return func(res http.ResponseWriter, req *http.Request, params matcher.PathParams) {
		switch fnType := reflect.TypeOf(customHandler); fnType.NumIn() {
		case 0:
			fn := customHandler.(func())
			fn()
		case 1:
			fn := customHandler.(func(res http.ResponseWriter))
			fn(res)
		case 2:
			fn := customHandler.(func(res http.ResponseWriter, req *http.Request))
			fn(res, req)
		case 3:
			fn := customHandler.(func(res http.ResponseWriter, req *http.Request, params matcher.PathParams))
			fn(res, req, params)
		default:
			panic("Custom handler got an unexpected count of input parameters.")
		}
	}
}
