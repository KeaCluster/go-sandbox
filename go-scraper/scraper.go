package main

import (
	"io"
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
		return "", err
	}
	defer res.Body.Close()

	// 200?
	if res.StatusCode != http.StatusOK {
		return "", err
	}

	// Parse
	title, err := extractTitle(res.Body)
	if err != nil {
		return "", err
	}
	return title, nil
}

func extractTitle(body io.Reader) (string, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return "", err
	}

	// Get title
	var f func(*html.Node) string
	f = func(n *html.Node) string {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			return n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			result := f(c)
			if result != "" {
				return result
			}
		}
		return ""
	}
	title := f(doc)
	// found
	if title == "" {
		return "", err
	}

	return strings.TrimSpace(title), nil
}
