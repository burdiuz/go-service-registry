package matcher

type PathMatch struct {
	Handler PathHandler
	Params  PathParams
}

// PathRegistry is a facade to path/matcher package
type PathRegistry struct {
	root *PathSegment
}

// PathRegistryNew creates new instance of PathRegistry
func PathRegistryNew() *PathRegistry {
	return &PathRegistry{root: PathSegmentNewRoot()}
}

// AddPath adds Path to comparison tree
func (r *PathRegistry) addPath(path *Path) error {
	return r.root.Insert(path, 0)
}

// Add adds new path to path comparison tree and puts handler on its end
func (r *PathRegistry) Add(pathStr string, handler PathHandler) error {
	path, err := PathNew(pathStr, handler)

	if err != nil {
		return err
	}

	return r.addPath(path)
}

// Get retrieves handler function for specified path or returns nil if nothing found
func (r *PathRegistry) Get(pathStr string) *PathMatch {
	parts := SplitURLPath(pathStr)
	path := r.root.Match(parts)

	if path == nil {
		return nil
	}

	var params PathParams = nil

	if path.HasParameters() {
		params = path.GetValuesFrom(parts)
	}

	return &PathMatch{Handler: path.Handler, Params: params}
}
