package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
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
	file := flag.String("f", "", "path to file with data")
	db := flag.Bool("db", false, "use database")
	flag.Parse()

	mux := defaultMux()

	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	fmt.Println("Starting the server on :8080")

	if *db {
		dbHandler, err := urlshort.DBHandler(mapHandler)
		if err != nil {
		} else if *db {
			panic(err)
		}
		http.ListenAndServe(":8080", dbHandler)
	} else if *file != "" {
		fileType := filepath.Ext(*file)

		switch fileType {
		case ".json":
			jsonData, err := os.ReadFile(*file)
			if err != nil {
				panic(err)
			}

			jsonHandler, err := urlshort.JSONHandler(jsonData, mapHandler)
			if err != nil {
				panic(err)
			}
			http.ListenAndServe(":8080", jsonHandler)
		case ".yaml":
			yamlData, err := os.ReadFile(*file)
			if err != nil {
				panic(err)
			}

			yamlHandler, err := urlshort.YAMLHandler(yamlData, mapHandler)
			if err != nil {
				panic(err)
			}
			http.ListenAndServe(":8080", yamlHandler)
		}
	} else {
		http.ListenAndServe(":8080", mapHandler)
	}
}
