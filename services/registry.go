package services

type ServiceRegistry struct {
	Services ServiceNameMap
}

func ServiceRegistryNew() *ServiceRegistry {
	registry := ServiceRegistry{Services: make(ServiceNameMap)}

	return &registry
}

func (sr *ServiceRegistry) Add(name, version, ipaddr, port string) *Service, error {
	s := ServiceNew(name, version, ipaddr, port)

	return s, sr.Services.Add(s)
}

func (sr *ServiceRegistry) Find(name, version string) []*Service {
	return sr.Services.Find(name, version)
}

func (sr *ServiceRegistry) GetExact(name, version string) []*Service {
	return sr.Services.Find(name, version)
}

func (sr *ServiceRegistry) Remove(name, version string) {
	return sr.Services.Remove§(name, version)
}
