package main

import (
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
	
	// Parse function return root node element from DOM tree, which has links to other node, which can be used to traver the whole DOM tree
	doc, err := html.Parse(strings.NewReader(htmlBody))

	if err != nil {
		return nil, err
	}
	
	URLs := []string{}

	for n := range doc.Descendants() {
		if n.DataAtom == atom.A {
			href := getHref(n)
			if href == "" {
				// skipping empty <a/> elements
			} else if strings.Contains(href, rawBaseURL){
				URLs = append(URLs, getHref(n))
			} else {
				URLs = append(URLs, rawBaseURL + getHref(n))
			}
		}
	}

	return URLs, nil
}


func getHref(n *html.Node) string {
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			return attr.Val
		}
	}
	return ""
}
