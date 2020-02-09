package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	httprouter "./http/router"
	matcher "./path/matcher"
)

func rootHandler(res http.ResponseWriter, req *http.Request, params matcher.PathParams) {
	message := "Somesing was sent."

	strings.Count(message, message)

	res.Write([]byte(message))
}

func main2() {

	registry := matcher.PathRegistryNew()

	registry.Add("/view/home", func(_ http.ResponseWriter, _ *http.Request, _ matcher.PathParams) {
		fmt.Println("Home handler worked!")
	})

	registry.Add("/view/:var/sort/Messages", func(_ http.ResponseWriter, _ *http.Request, params matcher.PathParams) {
		fmt.Printf("Variable handler worked: %v\n", params)
	})

	registry.Get("/view/home").Handler(nil, nil, nil)

	fmt.Printf("Handler: %v\n", registry.Get("/view/home/abc"))

	match := registry.Get("/view/global/sort/Messages")
	match.Handler(nil, nil, match.Params)

	path, err := matcher.PathNew("/view/home/sort/Messages", rootHandler)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Path: %v\n", path)

	path, err = matcher.PathNew("/Messages/home/view///sort", rootHandler)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Path: %v\n", path)
}

func main() {
	router := httprouter.New(nil)
	router.Route("/apple", nil).Get(func(w http.ResponseWriter, r *http.Request, v matcher.PathParams) {
		message := "Apple was called."

		w.Write([]byte(message))
	})

	router.Route("/some/:val", nil).Get(func(w http.ResponseWriter, r *http.Request, params matcher.PathParams) {
		message := "Some " + params["val"] + " was called."

		w.Write([]byte(message))
	})

	http.HandleFunc("/", router.GetHandler())

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("The end")
}
