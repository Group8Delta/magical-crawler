package test

import (
	"context"
	"encoding/json"
	"fmt"
	"magical-crwler/services/crawler"
	"runtime"
	"strconv"
	"testing"
	"time"
)

func TestMemStat(t *testing.T) {
	var memStatsStart, memStatsEnd runtime.MemStats

	runtime.ReadMemStats(&memStatsStart)
	a := []string{}
	for i := 0; i < 1_000_000; i++ {
		a = append(a, strconv.Itoa(i))
	}
	runtime.ReadMemStats(&memStatsEnd)
	fmt.Println("Mem: ", (memStatsEnd.Alloc-memStatsStart.Alloc)/(1024*1024), "MB")

}

func TestWorkerPool(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	wp := crawler.NewWorkerPool(url, 2, testDivarCrawler)

	wp.Start(ctx)
	results := wp.GetResults()
	errors := wp.GetErrors()
	fmt.Printf("results count:%d\n", len(results))
	fmt.Printf("errors count:%d\n", len(errors))

	for _, v := range errors {
		fmt.Println(v.Err.Error())
	}

}

func TestDivarCrawler(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"
	ctx, _ := context.WithTimeout(context.Background(), 200*time.Second)
	links, err := testDivarCrawler.CrawlAdsLinks(ctx, url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}

func TestSheypoorCrawler(t *testing.T) {
	url := "https://www.sheypoor.com/s/zanjan/houses-apartments-for-sale"
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	links, err := testSheypoorCrawler.CrawlAdsLinks(ctx, url)

	if err != nil {
		t.Fatalf("crawl adds links error:: %v", err)
	}

	t.Logf("crawled adds links: %v", links)
}

func TestCrawlDivarPageUrl(t *testing.T) {
	url := "https://divar.ir/v/زمین-فاز3-گلدشت/wZzogIfs"
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	ad, err := testDivarCrawler.CrawlPageUrl(ctx, url)

	if err != nil {
		t.Fatalf("crawl page url error:: %v", err)
	}
	if ad.Title == "" {
		t.Fatalf("error in crawling ad attributes")
	}

	j, err := json.Marshal(ad)
	if err != nil {
		t.Fatalf("error in marashal ad: %v", err)
	}
	t.Logf("result :\n %v", string(j))
}

func TestCrawlSheypoorPageUrl(t *testing.T) {
	url := "https://www.sheypoor.com/v/%D8%AE%D8%A7%D9%86%D9%87-%D8%A7%D8%AC%D8%A7%D8%B1%D9%87-%D8%A7-%D9%81%D9%86%D9%88%D8%B4-%D8%A2%D8%A8%D8%A7%D8%AF-445243769.html"
	ctx, _ := context.WithTimeout(context.Background(), 20*time.Second)
	ad, err := testSheypoorCrawler.CrawlPageUrl(ctx, url)

	if err != nil {
		t.Fatalf("crawl page url error: %v", err)
	}
	if ad.Title == "" {
		t.Fatalf("error in crawling ad attributes")
	}
	j, err := json.Marshal(ad)
	if err != nil {
		t.Fatalf("error in marashal ad: %v", err)
	}
	t.Logf("result :\n %v", string(j))
}

func TestSaveCrawlerData(t *testing.T) {
	err := crawler.SaveAdData(testRepo, &crawler.Ad{Title: "test1", Link: "http://test",
		PhotoUrl: "test_photo", SellerContact: "0914444444",
		Description: "test description", Price: 600000000,
		RentPrice: 5000000, City: "tehran", Lat: 5, Lon: 3,
		Neighborhood: "saadat abad", Size: 80, Bedrooms: 2,
		HasElevator: true, BuiltYear: 1400,
		IsApartment: true, CreationTime: time.Now()})

	if err != nil {
		t.Fatalf("error in saving crawler data:%v\n", err)
	}

}
