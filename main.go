package main

import (
	"net/http"

	httprouter "./http/router"
	matcher "./path/matcher"
)

func getServiceKey(name, version string) string {
	return name + "-|-" + version
}

type ServiceRegistry struct {
	services map[string]*Service
}

func (sr *ServiceRegistry) Register(name, version, ipaddr, port string) {

}

func (sr *ServiceRegistry) Find(name, version string) []*Service {

}

func (sr *ServiceRegistry) Remove(name, version string) {

}

func main() {
	router := httprouter.New(nil)

	router.Route("/registry/register/:name/:version/:port", func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {

	})
	router.Route("/registry/find/:name/:version", func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {

	})
	router.Route("/registry/unregister/:name/:version", func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {

	})
}
