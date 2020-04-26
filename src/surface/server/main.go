package main

import (
	"fmt"
	"log"
	"net/http"
	"surface"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()
		surfaceSvg := surface.Surface()
		w.Header().Set("Content-Type", "image/svg+xml")
		_, err := w.Write([]byte(surfaceSvg))
		if err != nil {
			log.Fatal("Error writing", err)
		}
	})
	http.HandleFunc("/counter", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Count: %d\n", count)
}

