package test

import (
	"magical-crwler/services"
	"testing"
)

func TestCrawl(t *testing.T) {
	url := "https://divar.ir/s/zanjan/buy-residential"
	services.CrawlHomeAds(url)
}
