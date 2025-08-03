package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"urlshort"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func main() {
	yaml := flag.String("f", "", "path to yaml file")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")

	if *yaml != "" {
		yamlFile, err := os.ReadFile(*yaml)
		if err != nil {
			panic(err)
		}

		yamlHandler, err := urlshort.YAMLHandler(yamlFile, mapHandler)
		if err != nil {
			panic(err)
		}

		http.ListenAndServe(":8080", yamlHandler)
	} else {
		http.ListenAndServe(":8080", mapHandler)
	}
}
