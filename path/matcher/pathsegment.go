package matcher

import "fmt"

// PathSegment represents a node in registered paths tree that is being matched against other paths to find endpint which will handle request
type PathSegment struct {
	Name     string
	Children map[string]*PathSegment
	Path     *Path
}

// PathSegmentNew creates a new PathSegment
func PathSegmentNew(name string, path *Path) *PathSegment {
	root := PathSegment{Name: name, Path: path}

	return &root
}

// PathSegmentNewRoot creates new PathSegment with empty name that can be used as root for PathSegment tree
func PathSegmentNewRoot() *PathSegment {
	return PathSegmentNew("", nil)
}

// IsLeaf checks if PathSegment is the last in tree branch and does not have Children
func (p *PathSegment) IsLeaf() bool {
	return p.Children == nil || len(p.Children) == 0
}

// IsEndpoint checks if PathSegment has Path attached to it, so it can handle requests
func (p *PathSegment) IsEndpoint() bool {
	return p.Path != nil
}

// IsVariable checks if PathSegment name is a variable
func (p *PathSegment) IsVariable() bool {
	return IsSegmentVar(p.Name)
}

// IsVariable checks if PathSegment name is a variable
func (p *PathSegment) String() string {
	return getSegmentIdentifier(p.Name)
}

// Insert adds more path segments to segment tree
func (p *PathSegment) Insert(path *Path, index int) error {
	isLast := index == path.Length()-1
	name := path.Path[index]
	key := getSegmentIdentifier(name)

	if p.Children == nil {
		p.Children = make(map[string]*PathSegment)
	}

	var node *PathSegment = p.Children[key]

	if node == nil {
		node = PathSegmentNew(name, nil)
		p.Children[key] = node
	}

	if isLast {
		if node.Path == nil {
			node.Path = path
		} else {
			return fmt.Errorf("Path %s already registered", path.String())
		}
	} else {
		return node.Insert(path, index+1)
	}

	return nil
}

// Match matches a list of path names against PathSegment tree to find an endpoint that fits it
func (p *PathSegment) Match(path []string) *Path {
	if p.IsLeaf() {
		return nil
	}

	return matchFrom(path, 0, p)
}

// MatchString matches a path against PathSegment tree to find an endpoint that fits it
func (p *PathSegment) MatchString(path string) *Path {
	parts := SplitURLPath(path)

	return p.Match(parts)
}
