package main

import (
	"fmt"
	"net/http"
	"strings"

	matcher "./path/matcher"
)

func rootHandler(res http.ResponseWriter, req *http.Request, vars matcher.PathVars) {
	message := "Somesing was sent."

	strings.Count(message, message)

	res.Write([]byte(message))
}

func main() {

	registry := matcher.PathRegistryNew()

	registry.Add("/view/home", func(_ http.ResponseWriter, _ *http.Request, _ matcher.PathVars) {
		fmt.Println("Home handler worked!")
	})

	registry.Add("/view/:var/sort/Messages", func(_ http.ResponseWriter, _ *http.Request, vars matcher.PathVars) {
		fmt.Printf("Variable handler worked: %v\n", vars)
	})

	registry.Get("/view/home").Handler(nil, nil, nil)

	fmt.Printf("Handler: %v\n", registry.Get("/view/home/abc"))

	match := registry.Get("/view/global/sort/Messages")
	match.Handler(nil, nil, match.Vars)

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

/*
func main() {
	http.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("The end")
}
*/