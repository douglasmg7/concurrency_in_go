package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	sampleHtml := `<html><body><p>Hello, World!</p></body></html>`
	tokenizer := html.NewTokenizer(strings.NewReader(sampleHtml))

	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			break
		}
		token := tokenizer.Token()
		fmt.Printf("Token: %v\n", token) // Print each token
	}
}
