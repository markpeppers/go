package link

import (
	"log"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseLinks(htmlDoc string) []Link {

	links := make([]Link, 0)
	r := strings.NewReader(htmlDoc)

	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			l := Link{}
			l.Href = n.Attr[0].Val
			l.Text = assembleText(n)
			links = append(links, l)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links
}

func assembleText(n *html.Node) string {
	return trimSurrounding(trimSpace(assembleTextStep(n, "")))
}

func assembleTextStep(n *html.Node, text string) string {
	if n.Type == html.TextNode {
		text = text + n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text = assembleTextStep(c, text)
	}
	return text
}

func trimSpace(s string) string {
	re := regexp.MustCompile(`\s+`)
	resultBytes := re.ReplaceAll([]byte(s), []byte(" "))
	return string(resultBytes)
}

func trimSurrounding(s string) string {
	re := regexp.MustCompile(`^ | $`)
	resultBytes := re.ReplaceAll([]byte(s), []byte(""))
	return string(resultBytes)
}
