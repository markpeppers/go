package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/markpeppers/link"
)

type Urlset struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	Urls    []Url    `xml:"url"`
}

type Url struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
}

type siteMapOpts struct {
	siteHost string
	scheme   string
}

func main() {
	siteUrl := "https://example.com/"
	if len(os.Args) < 2 {
		fmt.Printf("No site given. Using default %s\n", siteUrl)
	} else {
		siteUrl = os.Args[1]
	}

	resp, err := http.Get(siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	links := link.ParseLinks(string(body))

	siteUrlParsed, err := url.Parse(siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	siteHost := siteUrlParsed.Host
	scheme := siteUrlParsed.Scheme

	locations := collectLocations(
		siteMapOpts{
			siteHost: siteHost,
			scheme:   scheme,
		},
		links)

	fmt.Println(urlEncode(locations))
}

func collectLocations(opts siteMapOpts, links []link.Link) []string {
	visited := make(map[string]bool)

	locations := make([]string, 0)

	for _, l := range links {
		loc := l.Href
		// Ignore locations starting with "//"
		if len(loc) > 1 && loc[0:2] == "//" {
			continue
		}
		u, err := url.Parse(loc)
		if err != nil {
			log.Printf("error %s parsing %s, continuing...", err, loc)
		}
		// Ignore refs to external sites
		if len(u.Host) > 0 && u.Host != opts.siteHost {
			continue
		}
		// Add site host to relative URLs
		if len(u.Host) == 0 {
			loc = opts.scheme + "://" + opts.siteHost + u.Path
		} else {
			loc = opts.scheme + "://" + u.Host + u.Path
		}
		// Ignore repeated locations
		if visited[loc] {
			continue
		}
		visited[loc] = true
		locations = append(locations, loc)
	}

	return locations
}

func urlEncode(locs []string) string {
	doc := Urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		Urls:  []Url{},
	}
	for _, loc := range locs {
		doc.Urls = append(doc.Urls, Url{Loc: loc})
	}
	output, err := xml.MarshalIndent(doc, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	return xml.Header + string(output)
}
