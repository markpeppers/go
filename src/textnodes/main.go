package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "textnodes: %v", err)
		os.Exit(1)
	}
	strings := textnodes(nil, doc)
	for _, s := range strings {
		fmt.Println(s)
	}
}

func textnodes(textStrings []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		if len(strings.TrimSpace(n.Data)) > 0 {
			textStrings = append(textStrings, n.Data)
		}
	}
	// Skip script or style
	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		if n.FirstChild != nil {
			textStrings = textnodes(textStrings, n.FirstChild)
		}
	}
	if n.NextSibling != nil {
		textStrings = textnodes(textStrings, n.NextSibling)
	}
	return textStrings
}
