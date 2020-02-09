package matcher

import (
	"strings"
)

// SplitURLPath breaks path string into segments
func SplitURLPath(path string) []string {
	original := strings.Split(path, "/")
	filtered := make([]string, 0)

	/*
	  It is possible I need to distinguish between path/path and path/path/
	  So last segment could have an empty Name
	*/

	for _, part := range original {
		if part != "" {
			filtered = append(filtered, part)
		}
	}

	return filtered
}

// HasParamSegments checks if path contains parameter segments
func HasParamSegments(path string) bool {
	return strings.Contains(path, "/:")
}

// IsSegmentParam checks if segment string a variable declaration
func IsSegmentParam(name string) bool {
	return name[:1] == ":"
}

func getSegmentIdentifier(name string) string {
	if IsSegmentParam(name) {
		return ":"
	}

	return name
}

/*
	Rules for matching:
	If exact match for name found it is being taken, no parameters, even if it is not going to resolve
	For example, if these paths registered
	/path1/path2
	/path1/:var1/something/else
	And requested match is
	/path1/path2/something/else
	It will match to first one and finish unresolved because path2 is leaf node.
	So, if you registered a path with exact segment and a variable, the values for this variable can be anything but that exact name.
*/
func matchFrom(path []string, index int, parent *PathSegment) *Path {
	if parent.IsLeaf() {
		return nil
	}

	name := path[index]
	leaf := index == len(path)-1

	// first we take node by name match
	node := parent.Children[name]

	if node == nil {
		// if no match, then looking for a variable
		node = parent.Children[":"]

		if node == nil {
			// exit if no nodes found
			return nil
		}
	}

	if leaf {
		if node.IsEndpoint() {
			return node.Path
		}

		return nil
	}

	return matchFrom(path, index+1, node)
}

/*
func SplitURLPath(path string) ([]string, error) {
	var rgx *regexp.Regexp
	rgx, err := regexp.Compile("[^/]+")

	if err != nil {
		return nil, err
	}

	return rgx.FindAllString(path, 0xFF), nil
}
*/
