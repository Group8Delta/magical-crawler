package crawler

import (
	"errors"
	"magical-crwler/config"
)

type CrawlerType string

const (
	DivarCrawlerType    CrawlerType = "divar_crawler"
	SheypoorCrawlerType CrawlerType = "sheypoor_crawler"
)
const numberOfCrawlerWorkers = 2

var CrawlerTypes []CrawlerType = []CrawlerType{
	DivarCrawlerType,
	SheypoorCrawlerType,
}

type CrawlerInterface interface {
	CrawlAdsLinks(searchUrl string) ([]string, error)
	CrawlPageUrl(pageUrl string) (*Ad, error)
	RunCrawler()
}

func New(crawlerType CrawlerType, config *config.Config, maxDeepth int) (CrawlerInterface, error) {
	switch crawlerType {
	case DivarCrawlerType:
		return &DivarCrawler{config: config, maxDeepth: maxDeepth}, nil
	case SheypoorCrawlerType:
		return &SheypoorCrawler{config: config, maxDeepth: maxDeepth}, nil
	default:
		return nil, errors.New("invalid crawler type")
	}
}
