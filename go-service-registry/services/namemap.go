package services

type ServiceNameMap map[string]ServiceVersionMap

func (snm ServiceNameMap) Add(s *Service) error {
	versions := snm[s.Name]

	if versions == nil {
		versions = make(ServiceVersionMap)
		snm[s.Name] = versions
	}

	return versions.Add(s)
}

func (snm ServiceNameMap) Versions(name string) ServiceVersionMap {
	return snm[name]
}

func (snm ServiceNameMap) Has(name, version string) bool {
	versions := snm.Versions(name)

	if versions == nil {
		return false
	}

	return versions.Has(version)
}

func (snm ServiceNameMap) HasExact(name, version string) bool {
	versions := snm.Versions(name)

	if versions == nil {
		return false
	}

	return versions.HasExact(version)
}

func (snm ServiceNameMap) Filter(name string, filter func(s *Service) bool) {
	versions := snm.Versions(name)

	if versions != nil {
		versions.Filter(filter)
	}
}

func (snm ServiceNameMap) FilterAll(filter func(s *Service) bool) {
	for _, versions := range snm {
		versions.Filter(filter)
	}
}

func (snm ServiceNameMap) Find(name, version string, filter func(s *Service) bool) ServiceList {
	versions := snm.Versions(name)

	if versions == nil {
		return ServiceList{}
	}

	return versions.Find(version, filter)
}

func (snm ServiceNameMap) Get(name, version string, filter func(s *Service) bool) *Service {
	versions := snm.Versions(name)

	if versions == nil {
		return nil
	}

	return versions.Get(version, filter)
}

func (snm ServiceNameMap) GetExact(name, version string) *Service {
	versions := snm.Versions(name)

	if versions == nil {
		return nil
	}

	return versions.GetExact(version)
}

func (snm ServiceNameMap) Remove(name, version string) {
	versions := snm.Versions(name)

	if versions != nil {
		versions.Remove(version)
	}
}
