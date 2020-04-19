package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const key string = "6d8847e8"
const urlbase string = "https://www.omdbapi.com/"

const templ = `*** {{.Title}} ***
Runtime:  {{.Runtime | runTime}}
Director: {{.Director}}
Age:      {{.Released | yearsAgo}} years
`

// OmdbMovie represents a movie
type OmdbMovie struct {
	Title    string
	Poster   string
	Runtime  string
	Director string
	Released string
}

func yearsAgo(date string) int {
	t, _ := time.Parse("02 Jan 2006", date)
	return int(time.Since(t).Hours() / 24 / 365)
}

func runTime(in string) string {
	mins, _ := strconv.Atoi(strings.Split(in, " ")[0])
	hours := int(mins / 60)
	return fmt.Sprintf("%dh %dm", hours, mins-hours*60)
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

	movieReport := template.Must(template.New("movie").
		Funcs(template.FuncMap{
			"yearsAgo": yearsAgo,
			"runTime":  runTime,
		}).
		Parse(templ))
	err = movieReport.Execute(os.Stdout, result)
	if err != nil {
		log.Fatal(err)
	}
}
