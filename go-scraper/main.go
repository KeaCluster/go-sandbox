package main

import "fmt"

func main() {
	urlList := []string{
		"www.wikileaks.org",
		"www.wikipedia.org",
	}

	workers := 5
	urls := make(chan string, len(urlList))
	results := make(chan ScrapingRes, len(urlList))

	for i := 0; i < workers; i++ {
		go worker(urls, results)
	}

	for _, url := range urlList {
		urls <- url
	}
	close(urls)

	for i := 0; i < len(urlList); i++ {
		result := <-results
		if result.Err != nil {
			fmt.Printf("Error scraping %s: %v\n", result.Url, result.Err)
		} else {
			fmt.Printf("Scraped %s: %s\n", result.Url, result.Title)
		}
	}
}
