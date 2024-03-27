package main

func worker(urls <-chan string, results chan<- ScrapingRes) {
	for url := range urls {
		title, err := scrape(url)
		results <- ScrapingRes{Url: url, Title: title, Err: err}
	}
}

func scrape(url string) (string, error) {

}
