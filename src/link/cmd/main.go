package main

import (
	"fmt"
	"io/ioutil"
	"link"
	"log"
	"os"
)

func main() {
	var filename string
	if len(os.Args) < 2 {
		filename = "test/ex2.html"
	} else {
		filename = os.Args[1]
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	links := link.ParseLinks(string(b))

	for _, l := range links {
		fmt.Println("Link")
		fmt.Printf("  Href: \"%s\"\n", l.Href)
		fmt.Printf("  Text: \"%s\"\n", l.Text)
	}
}
