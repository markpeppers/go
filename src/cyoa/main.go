package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

const templ = `<h1>{{.Title}}</h1>
{{range .Story}}
<p>{{.}}</p>
{{end}}
{{range .Options}}
<p><a href="{{.Arc}}">{{.Text}}</a></p>
{{end}}
`

var storyFile string

type Adventure map[string]Arc

type Arc struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func init() {
	flag.StringVar(&storyFile, "i", "gopher.json", "json file containing the adventure")
	flag.Parse()
}

func main() {
	storyJson, err := ioutil.ReadFile(storyFile)
	if err != nil {
		log.Fatal(err)
	}

	var adventure Adventure

	err = json.NewDecoder(bytes.NewReader(storyJson)).Decode(&adventure)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", adventure))
}

func (a Adventure) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := r.URL.RequestURI()
	req = strings.ReplaceAll(req, "/", "")
	if len(req) == 0 {
		req = "intro"
	}
	arc, ok := a[req]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "You look to be lost, young gopher.")
		return
	}

	t := template.Must(template.New("arc").Parse(templ))
	t.Execute(w, arc)
}
