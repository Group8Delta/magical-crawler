package test

import (
	"fmt"
	"magical-crwler/services/crawler"
	"testing"
)

func TestWorkerPool(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"

	wp := crawler.NewWorkerPool(url, 2, testDivarCrawler)

	wp.Start()
	results := wp.GetResults()
	errors := wp.GetErrors()

	fmt.Printf("results count:%d\n", len(results))
	fmt.Printf("errors count:%d\n", len(errors))

	fmt.Println(errors)

}

func TestDivarCrawler(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"

	links, err := testDivarCrawler.CrawlAdsLinks(url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}

func TestSheypoorCrawler(t *testing.T) {
	url := "https://www.sheypoor.com/s/zanjan/houses-apartments-for-sale"

	links, err := testSheypoorCrawler.CrawlAdsLinks(url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}

func TestCrawlDivarPageUrl(t *testing.T) {
	url := "https://divar.ir/v/%D8%A2%D9%BE%D8%A7%D8%B1%D8%AA%D9%85%D8%A7%D9%86-%DB%B6%DB%B8%D9%85%D8%AA%D8%B1%DB%8C-%D8%AC%D9%86%D9%88%D8%A8%DB%8C-%D8%B2%DB%8C%D8%A8%D8%A7-%D8%A8%D8%A7-%D9%85%D8%AA%D8%B1%DB%8C%D8%A7%D9%84/wZLIKFFJ"

	ad, err := testDivarCrawler.CrawlPageUrl(url)

	if err != nil {
		t.Fatalf("crawl page url error:: %v", err)
	}

	t.Logf("result : %v", ad)
}

func TestCrawlSheypoorPageUrl(t *testing.T) {
	url := "https://www.sheypoor.com/v/%D8%AE%D8%A7%D9%86%D9%87-%D8%A7%D8%AC%D8%A7%D8%B1%D9%87-%D8%A7-%D9%81%D9%86%D9%88%D8%B4-%D8%A2%D8%A8%D8%A7%D8%AF-445243769.html"

	ad, err := testSheypoorCrawler.CrawlPageUrl(url)

	if err != nil {
		t.Fatalf("crawl page url error: %v", err)
	}

	t.Logf("result : %v", ad)
}
