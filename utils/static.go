package utils

import (
	"net/http"

	matcher "../path/matcher"
)

// HandleStatic maps request path to directory content and returns file content if found
func HandleStatic(rootDir string) matcher.PathHandler {
	handler := http.FileServer(http.Dir(rootDir))

	return func(res http.ResponseWriter, req *http.Request, params matcher.PathParams) {
		handler.ServeHTTP(res, req)
	}
}

/*HandleStaticPartial works just like HandleStatic except it cuts rootHTTPPath content from
  beginning of HTTP path before matching file system content
*/
func HandleStaticPartial(rootDir string, rootHTTPPath string) matcher.PathHandler {
	handler := http.StripPrefix(rootHTTPPath, http.FileServer(http.Dir(rootDir)))

	return func(res http.ResponseWriter, req *http.Request, params matcher.PathParams) {
		handler.ServeHTTP(res, req)
	}
}
