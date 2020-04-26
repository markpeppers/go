package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "need at least one arg")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %s: %v\n", url, err)
			continue
		}
		fmt.Printf("%s => %s, %d bytes\n", url, filename, n)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		fmt.Println("closing file", f.Name())
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
