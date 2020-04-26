// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// Copyright © 2020 Mark A. Peppers.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Printf(" %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Print("/>\n")
			return
		}
		if hasTextChild(n) {
			fmt.Print(">")
		} else {
			fmt.Print(">\n")
		}
		depth++
		return
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!-- %s -->\n", depth*2, "", n.Data)
		return
	}
	if n.Type == html.TextNode {
		fmt.Printf("%s", strings.TrimSpace(n.Data))
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		if n.FirstChild != nil {
			depth--
			if hasTextChild(n) {
				fmt.Printf("</%s>\n", n.Data)
				return
			}
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func hasTextChild(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && len(strings.TrimSpace(c.Data)) > 0 {
			return true
		}
	}
	return false
}

//!-startend
