package services

import (
	"fmt"

	semver "golang.org/x/tools/internal/semver"
)

// TODO make it an array of services?
type ServiceVersionMap map[string]*Service

func (svm ServiceVersionMap) Add(s *Service) {
	if svm[s.Version] != nil {
		return fmt.Errorf("Service %q version %q is already registered", s.Name, s.Version)
	}

	svm[s.Version] = s
	return nil
}

func (svm ServiceVersionMap) Has(version string) bool {
	for sver, _ := range svm {
		if semver.Compare(sver, version) == 0 {
			return true
		}
	}

	return false
}

func (svm ServiceVersionMap) HasExact(version string) bool {
	return svm[version] != nil
}

func (svm ServiceVersionMap) Find(version string) []*Service {
	list := make([]*Service, 0)

	for sver, s := range svm {
		if semver.Compare(sver, version) == 0 {
			append(list, s)
		}
	}

	return list
}

func (svm ServiceVersionMap) Get(version string) *Service {
	for sver, s := range svm {
		if semver.Compare(sver, version) == 0 {
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
