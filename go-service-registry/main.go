package main

import (
	"fmt"
	"log"
	"net/http"

	httprouter "../go-service-router"
	matcher "../go-url-path-matcher"
	services "./services"
	utils "./utils"
)

func main() {
	registry := services.NewServiceRegistry()
	router := httprouter.New(nil)

	/*
		It must have port value to specify the port service is listening to,
		not from where request was sent
	*/
	router.Route("/register/:name/:version/:port", nil).Post(func(
		w http.ResponseWriter,
		r *http.Request,
		params matcher.PathParams) {
		// TODO Combine Register & Refresh
		name := params["name"]
		version := params["version"]
		addr := utils.GetRemoteAddrIp(r.RemoteAddr)
		port := params["port"]

		s, err := registry.Add(name, version, addr, port)

		fmt.Printf("Registered: %v\n", s)

		if err != nil {
			log.Fatalf("%v\n", err)
		}

		err = s.WriteHTTP(w)

		if err != nil {
			log.Fatalf("Getting services error: %v", err)
		}

	})

	router.Route("/register/:name/:version", func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {
		name := params["name"]
		version := params["version"]

		list := registry.Find(name, version)
		err := list.WriteHTTP(w)

		if err != nil {
			log.Fatalf("Getting services error: %v", err)
		}
	}).Put(func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {
		name := params["name"]
		version := params["version"]
		addr := r.RemoteAddr

		s, err := registry.Refresh(name, version, addr)

		fmt.Printf("Refresh: %v\n", s)

		if err != nil {
			log.Fatalf("%v\n", err)
		}

		err = s.WriteHTTP(w)

		if err != nil {
			log.Fatalf("Getting services error: %v", err)
		}
	}).Delete(func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {
		name := params["name"]
		version := params["version"]

		fmt.Printf("Before Remove: %v\n", registry.GetExact(name, version))

		registry.Remove(name, version)

		fmt.Printf("After Remove: %v\n", registry.GetExact(name, version))

		w.Write([]byte("{ message: \"Service removed\"}"))
	})

	http.HandleFunc("/", router.GetHandler())

	/*
		expired := services.NewExpiredChecks(registry, 30*time.Second)
		expired.Start(3)
	*/

	/*
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Hello World: " + r.URL.Path))
		})
	*/

	log.Fatal(http.ListenAndServe(":8080", nil))
}
