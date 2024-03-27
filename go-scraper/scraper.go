package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func worker(urls <-chan string, results chan<- ScrapingRes) {
	for url := range urls {
		title, err := scrape(url)
		results <- ScrapingRes{Url: url, Title: title, Err: err}
	}
}

func scrape(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("Error scraping %s: %v", url, err)
	}
	defer res.Body.Close()

	// 200?
	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Unexpected status code %d for %s", res.StatusCode, url)
	}

	// Parse
	doc, err := html.Parse(res.Body)
	if err != nil {
		return "", fmt.Errorf("Error parsing HTML for %s: %v", url, err)
	}

	// Get title
	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" {
			if n.FirstChild != nil {
				title = n.FirstChild.Data
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	// found
	if title != "" {
		return strings.TrimSpace(title), nil
	}

	return "", fmt.Errorf("No title found for %s, ", url)
}
