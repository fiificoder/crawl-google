package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func main() {
	// Set the search query
	query := url.Values{"q": {"golang web scraper"}}

	// Make an HTTP GET request to the Google search URL
	res, err := http.Get("https://www.google.com/search?" + query.Encode())
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	// Parse the response as HTML
	doc, err := html.Parse(res.Body)
	if err != nil {
		panic(err)
	}

	// Use the `Visit` function to iterate through all the nodes in the HTML document
	var f func(*html.Node)
	f = func(n *html.Node) {
		// Check if the node is an anchor tag
		if n.Type == html.ElementNode && n.Data == "a" {
			// Iterate through the anchor tag's attributes
			for _, attr := range n.Attr {
				// Check if the attribute is the "href" attribute
				if attr.Key == "href" {
					// Print the value of the "href" attribute
					fmt.Println(attr.Val)
					break
				}
			}
		}
		// Recursively visit all the child nodes
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
}
