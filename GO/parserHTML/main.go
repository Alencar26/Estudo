package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	r := strings.NewReader(`
<html>
  <body>
    <h1>Olá</h1>
    <p>Texto</p>
  </body>
</html>
`)

	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}

	Walk(doc)
}

func Walk(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Println("Tag: ", n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Walk(c)
	}
}
