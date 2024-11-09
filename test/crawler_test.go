package test

import (
	"fmt"
	"magical-crwler/services/crawler"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"
	c, err := crawler.New(crawler.DivarCrawlerType)
	wp := crawler.NewWorkerPool(url, 2, c)
	if err != nil {
		t.Fatalf("initial crawler error: %v", err)
	}
	wp.Start()
	results := wp.GetResults()
	fmt.Println(len(results))
}

func TestDivarCrawler(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"
	c, err := crawler.New(crawler.DivarCrawlerType)
	if err != nil {
		t.Fatalf("initial crawler error: %v", err)
	}
	links, err := c.CrawlAdsLinks(url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}

func TestSheypoorCrawler(t *testing.T) {
	url := "https://www.sheypoor.com/s/zanjan/houses-apartments-for-sale"

	c, err := crawler.New(crawler.SheypoorCrawlerType)
	if err != nil {
		t.Fatalf("initial crawler error: %v", err)
	}
	links, err := c.CrawlAdsLinks(url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}
