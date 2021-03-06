package matcher

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	utils "./utils"
)

// PathParams is an alias of a map that contains values of path parameters
type PathParams map[string]string

// PathHandler is an alias for HTTP request handler function
type PathHandler func(w http.ResponseWriter, r *http.Request, params PathParams)

/*PathParamsHandler is an alias for HTTP request handler function with additional params argument
type PathParamsHandler func(w http.ResponseWriter, r *http.Request, params PathParams)
*/

func splitPath(path string) ([]string, error) {
	if strings.Contains(path, "//") {
		return nil, errors.New("Path cannot contain empty segments")
	}

	return utils.SplitURLPath(path), nil
}

// Path structure that represents a endpoint with path and string
type Path struct {
	Path       []string
	parameters map[int]string
	Handler    PathHandler
}

// NewPath creates new Path struct instance
func NewPath(path string, handler PathHandler) (*Path, error) {
	if handler == nil {
		return nil, errors.New("Path must have a valid handler")
	}

	parts, err := splitPath(path)

	if err != nil {
		return nil, err
	}

	var parameters map[int]string = nil
	names := make(map[string]bool)

	if HasParamSegments(path) {
		parameters = make(map[int]string)

		for index, part := range parts {
			if IsSegmentParam(part) {
				name := part[1:]

				if names[name] {
					return nil, fmt.Errorf("Path parameter %q used twice in %q", name, path)
				}

				names[name] = true
				parameters[index] = name
			}
		}
	}

	return &Path{Path: parts, parameters: parameters, Handler: handler}, nil
}

// Length returns now many segments are in the path
func (p *Path) Length() int {
	return len(p.Path)
}

// HasParameters returns true if Path has parameter segments
func (p *Path) HasParameters() bool {
	return p.parameters != nil
}

// GetValuesFrom returns a map with variable name from path and its value from passed path
func (p *Path) GetValuesFrom(path []string) PathParams {
	if !p.HasParameters() {
		return nil
	}

	values := make(map[string]string)

	for index, name := range p.parameters {
		values[name] = path[index]
	}

	return values
}

// GetValuesFromString returns a map with variable name from path and its value from passed path
func (p *Path) GetValuesFromString(path string) map[string]string {
	return p.GetValuesFrom(utils.SplitURLPath(path))
}

func (p *Path) String() string {
	return strings.Join(p.Path, "/")
}
