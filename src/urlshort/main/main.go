package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"urlshort"
)

var yamlFilePath string

func init() {
	flag.StringVar(&yamlFilePath, "y", "", "(optional) yaml file containing paths and urls")
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	// http.ListenAndServe(":8080", mapHandler)
	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	if len(yamlFilePath) > 0 {
		yamlBytes, err := ioutil.ReadFile(yamlFilePath)
		if err == nil {
			yaml = string(yamlBytes)
			fmt.Println("Read yaml from", yamlFilePath)
		}
	}
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	json := `
[
	{
		"path": "/c",
		"url": "https://cnn.com"
	},
	{
		"path": "/r",
		"url": "https://www.reddit.com/"
	}
]
`
	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
