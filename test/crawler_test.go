package test

import( 
	"testing"
	"magical-crwler/services"
)


func TestCrawl(t *testing.T){
	url:="https://divar.ir/s/zanjan/buy-residential"
	services.CrawlHomeAds(url)
}