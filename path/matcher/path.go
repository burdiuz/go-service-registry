package matcher

import (
	"errors"
	"net/http"
	"strings"
)

// PathVars is an alias of a map that contains values of path variables
type PathVars map[string]string

// PathHandler is an alias for HTTP request handler function
type PathHandler func(w http.ResponseWriter, r *http.Request, vars PathVars)

func splitPath(path string) ([]string, error) {
	if strings.Contains(path, "//") {
		return nil, errors.New("Path cannot contain empty segments")
	}

	return SplitURLPath(path), nil
}

// Path structure that represents a endpoint with path and string
type Path struct {
	Path      []string
	variables map[int]string
	Handler   PathHandler
}

// PathNew creates new Path struct instance
func PathNew(path string, handler PathHandler) (*Path, error) {
	if handler == nil {
		return nil, errors.New("Path must have a valid handler")
	}

	parts, err := splitPath(path)

	if err != nil {
		return nil, err
	}

	variables := make(map[int]string)

	for index, part := range parts {
		if IsSegmentVar(part) {
			variables[index] = part[1:]
		}
	}

	return &Path{Path: parts, variables: variables, Handler: handler}, nil
}

// Length returns now many segments are in the path
func (p *Path) Length() int {
	return len(p.Path)
}

// GetValuesFrom returns a map with variable name from path and its value from passed path
func (p *Path) GetValuesFrom(path []string) PathVars {
	values := make(map[string]string)

	for index, name := range p.variables {
		values[name] = path[index]
	}

	return values
}

// GetValuesFromString returns a map with variable name from path and its value from passed path
func (p *Path) GetValuesFromString(path string) map[string]string {
	return p.GetValuesFrom(SplitURLPath(path))
}

func (p *Path) String() string {
	return strings.Join(p.Path, "/")
}
