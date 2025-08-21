package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error){
	
	// Parse function return root node element from DOM tree, which has links to other node, which can be used to traver the whole DOM tree
	doc, err := html.Parse(strings.NewReader(htmlBody)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse body HTML %v",err)
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("Couldn't parse baseURL: %v", err)
	}
	
	URLs := []string{}

	for n := range doc.Descendants() {
		if n.DataAtom == atom.A {
			unparsedhref := getHref(n)
			
			href, err := url.Parse(unparsedhref)
			if err != nil {
				fmt.Printf("Couldn't parse href '%v': %v/n", unparsedhref, err)
				continue
			}
			
			if href != "" { // Skippin empty link elements
				resolvedURL := baseURL.ResolveRefernce(href)
				URLs = append(URLs, resolvedURL.String())
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
	return "" //empty href
}
