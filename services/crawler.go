package services

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"math/rand"

	"github.com/gocolly/colly"
)

var userAgents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Firefox/79.0",
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:79.0) Gecko/20100101 Firefox/79.0",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Safari/605.1.15",
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:80.0) Gecko/20100101 Firefox/80.0",
	"Mozilla/5.0 (iPhone; CPU iPhone OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15A372 Safari/604.1",
	"Mozilla/5.0 (Linux; Android 10; SM-G975F) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Mobile Safari/537.36",
	"Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.121 Safari/537.36",
	"Mozilla/5.0 (iPad; CPU OS 14_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/14.0 Mobile/15A5341f Safari/604.1",
}

func CrawlDivarAds(url string) {
	// Instantiate main collector for the initial page
	mainCollector := colly.NewCollector(
		colly.AllowedDomains("www.divar.ir", "divar.ir"),
	)

	mainCollector.Limit(&colly.LimitRule{
		DomainGlob:  "*divar.ir*", // Applies to all subdomains of divar.ir
		Delay:       200 * time.Millisecond,
		RandomDelay: 500 * time.Millisecond, // Adds up to 0.5 seconds of randomness
	})
	mainCollector.SetRequestTimeout(30 * time.Second) // Set a timeout for requests

	mainCollector.OnError(func(r *colly.Response, err error) {
		if r.StatusCode == http.StatusTooManyRequests { // Too Many Requests
			fmt.Println("Rate limit hit, waiting and retrying...")
			time.Sleep(2 * time.Second) // Wait before retrying
			r.Request.Retry()
		} else {
			fmt.Printf("Request to %s failed: %v\n", r.Request.URL, err)
		}
	})
	subCollector := mainCollector.Clone()
	// On every <a> element with the class "kt-post-card__action", visit the link
	mainCollector.OnHTML("a.kt-post-card__action", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		subLink := fmt.Sprintf("https://divar.ir%s", link)
		subCollector.Visit(subLink)
	})

	// Logging request URLs for mainCollector
	mainCollector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
		fmt.Println("Visiting main page:", r.URL.String())
	})

	// Extract and print information from the ad pages using subCollector
	subCollector.OnHTML("div.kt-page-title", func(e *colly.HTMLElement) {
		fmt.Println("Ad Title:", e.Text)
	})

	// Error handling for both collectors
	mainCollector.OnError(func(r *colly.Response, err error) {
		log.Println("Error on main page:", r.Request.URL, err)
	})
	subCollector.OnError(func(r *colly.Response, err error) {
		log.Println("Error on ad page:", r.Request.URL, err)
	})

	// Start the crawl with the main URL
	if err := mainCollector.Visit(url); err != nil {
		log.Fatal("Failed to visit URL:", err)
	}
}
