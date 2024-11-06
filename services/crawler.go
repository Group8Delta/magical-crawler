package services

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
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
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	max_deepth := 40
	deepth := 0

	var lastHeight, newHeight int64
	var allHTMLContent strings.Builder

	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("load page deepth : ",deepth)
		err = chromedp.Run(ctx,
			chromedp.Evaluate(`document.body.scrollHeight`, &newHeight),
		)

		if err != nil {
			log.Fatal(err)
		}

		if newHeight == lastHeight {
			fmt.Println("No more content to load.")
			break
		}

		lastHeight = newHeight

		var buttonExists bool
		err = chromedp.Run(ctx,
			// This evaluates whether the button is visible on the page
			chromedp.Evaluate(`document.querySelector('.post-list__load-more-btn-be092') !== null`, &buttonExists),
		)
		if err != nil {
			log.Fatal(err)
		}
		
		if buttonExists {
			err = chromedp.Run(ctx,
				chromedp.Click(".post-list__load-more-btn-be092", chromedp.ByQuery),
				chromedp.Sleep(500*time.Millisecond), // Wait for ads to load
			)
			fmt.Println("click load more")
		} else {
			err = chromedp.Run(ctx,
				chromedp.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`, nil),
				chromedp.Sleep(500*time.Millisecond),
			)
		}

		if err != nil {
			log.Fatal(err)
		}

		var html string
		err = chromedp.Run(ctx, chromedp.OuterHTML("html", &html))
		if err != nil {
			log.Fatal(err)
		}

		allHTMLContent.WriteString(html)
		deepth++
		if deepth == max_deepth {
			break
		}
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(allHTMLContent.String()))
	if err != nil {
		log.Fatal(err)
	}

	count := 0
	doc.Find("a.kt-post-card__action").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println("Link: ", href)
			count++
		} else {
			fmt.Println("No href found in h1 link.")
		}
	})
	fmt.Println("total :", count)

}

func CrawlDivarAds2(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, 100*time.Second)
	defer cancel()

	var allHTMLContent strings.Builder

	// Navigate to the URL
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
	)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Click the "Load More" button if it exists
		err = chromedp.Run(ctx,
			chromedp.Click(`button[data-virtuoso-scroller="true"]`, chromedp.NodeVisible),
			chromedp.Sleep(2*time.Second), // Wait for new ads to load
		)
		if err != nil {
			// If the button is no longer found, break the loop
			fmt.Println("No more 'Load More' button or failed to click.")
			break
		}

		// Capture the updated HTML content
		var html string
		err = chromedp.Run(ctx, chromedp.OuterHTML("html", &html))
		if err != nil {
			log.Fatal(err)
		}

		allHTMLContent.WriteString(html)
	}

	// Parse the combined HTML content with goquery
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(allHTMLContent.String()))
	if err != nil {
		log.Fatal(err)
	}

	// Extract the ad links
	count := 0
	doc.Find("a.kt-post-card__action").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			fmt.Println("Link: ", href)
			count++
		} else {
			fmt.Println("No href found in link.")
		}
	})
	fmt.Println("Total ads found:", count)
}
