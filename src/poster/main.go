package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strings"
)

const key string = "6d8847e8"
const urlbase string = "https://www.omdbapi.com/"

// OmdbMovie represents a movie
type OmdbMovie struct {
	Title  string
	Poster string
}

func main() {
	var terms []string
	for i := 1; i < len(os.Args); i++ {
		terms = append(terms, os.Args[i])
	}
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(urlbase + "?t=" + q + "&apikey=" + key)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var result OmdbMovie
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}

	// Get the extension, ie "jpeg"
	toks := strings.Split(result.Poster, ".")
	ext := toks[len(toks)-1]

	poster, err := http.Get(result.Poster)
	if err != nil {
		log.Fatal(err)
	}
	defer poster.Body.Close()

	posterFilename := strings.Join(terms, "+") + "." + ext
	f, err := os.Create(posterFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	io.Copy(w, poster.Body)

	cmd := exec.Command("open", posterFilename)
	cmd.Run()
}
