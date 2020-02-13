package services

import (
	"fmt"

	semver "golang.org/x/mod/semver"
)

// TODO make it a map of arrays of services?
type ServiceVersionMap map[string]*Service

func (svm ServiceVersionMap) Add(s *Service) error {
	if svm[s.Version] != nil {
		return fmt.Errorf("Service %q version %q is already registered", s.Name, s.Version)
	}

	svm[s.Version] = s
	return nil
}

func (svm ServiceVersionMap) Has(version string) bool {
	for sver := range svm {
		if semver.Compare(sver, version) == 0 {
			return true
		}
	}

	return false
}

func (svm ServiceVersionMap) HasExact(version string) bool {
	return svm[version] != nil
}

func (svm ServiceVersionMap) Filter(filter func(s *Service) bool) {
	for sver, s := range svm {
		if !filter(s) {
			delete(svm, sver)
		}
	}
}

func (svm ServiceVersionMap) Find(version string, filter func(s *Service) bool) ServiceList {
	list := make(ServiceList, 0)

	for sver, s := range svm {
		if semver.Compare(sver, version) == 0 && (filter == nil || filter(s)) {
			list = append(list, s)
		}
	}

	return list
}

func (svm ServiceVersionMap) Get(version string, filter func(s *Service) bool) *Service {
	for sver, s := range svm {
		if semver.Compare(sver, version) == 0 && (filter == nil || filter(s)) {
			return s
		}
	}

	return nil
}

func (svm ServiceVersionMap) GetExact(version string) *Service {
	return svm[version]
}

func (svm ServiceVersionMap) Remove(version string) {
	delete(svm, version)
}
