package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	sampleHtml := `<html><body><p class="as">Hello, World!</p><p>100</p><p>200</p></body></html>`
	doc, err := html.Parse(strings.NewReader(sampleHtml))
	if err != nil {
		panic(err)
	}
	traverse(doc)
}

func traverse(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("Element: %s\n", n.Data) // Print element names
	}
	if n.Type == html.TextNode {
		fmt.Printf("\t%s\n", n.Data) // Print element names
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c)
	}
}
