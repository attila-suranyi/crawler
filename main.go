package main

import (
	"net/http"
	"net/url"
)

//A web crawler which crawls a large amount of urls concurrently.
//Urls are loaded from sitemaps. Determine crawlable urls using the robots.txt.
//Crawling will be done using the worker pool pattern, so configuring concurrency can be done. This is important to not exhaust our resources, and not to get throttled or banned from server side.
//Rate limit the crawler, some delay should be added between requests.
//Add retry logic with exponential backoff using context deadline, log and save failed urls.
//The response will be processed, in whatever way necessary, so storing the page response for a time should be considered memory wise.

func main() {

}

type Crawler struct {
	client *http.Client
	wp WorkerPool
}

func initCrawler(cfg Config) *Crawler {
	return &Crawler{
		client: &http.Client{},
	}
}

func (cr *Crawler) getUrls(domain string) []url.URL {
	//Load sitemaps, parse for urls

	//Load robots.txt

	//Filter urls according to robots.txt
	return nil
}