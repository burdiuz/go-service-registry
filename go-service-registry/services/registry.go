package services

import (
	"fmt"
	"time"
)

type ServiceRegistry struct {
	Services ServiceNameMap
}

func NewServiceRegistry() *ServiceRegistry {
	registry := ServiceRegistry{make(ServiceNameMap)}

	return &registry
}

func (sr *ServiceRegistry) Add(name, version, remoteAddr, remotePort string) (*Service, error) {
	s := NewService(name, version, remoteAddr, remotePort)

	err := sr.Services.Add(s)

	if err != nil {
		return nil, err
	}

	return s, nil
}

func (sr *ServiceRegistry) Refresh(name, version, remoteAddr string) (*Service, error) {
	s := sr.Services.GetExact(name, version)

	if s == nil {
		return nil, fmt.Errorf("Service %q of %q version is not registered", name, version)
	} else if s.RemoteAddr == remoteAddr {
		return nil, fmt.Errorf("Service %q of %q version is registered but remote differs", name, version)
	}

	s.Touch()

	return s, nil
}

func (sr *ServiceRegistry) Find(name, version string) ServiceList {
	return sr.Services.Find(name, version, nil)
}

func (sr *ServiceRegistry) GetExact(name, version string) *Service {
	return sr.Services.GetExact(name, version)
}

func (sr *ServiceRegistry) Remove(name, version string) {
	sr.Services.Remove(name, version)
}

func (sr *ServiceRegistry) FilterExpired(ttl time.Duration) {
	sr.Services.FilterAll(func(s *Service) bool {
		return !s.IsExpired(ttl)
	})
}
