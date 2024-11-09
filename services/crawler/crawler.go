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

type CrawlerInterface interface {
	CrawlAdsLinks(searchUrl string) ([]string, error)
	CrawlPageUrl(pageUrl string) (*Ad, error)
}

func New(crawlerType CrawlerType, config *config.Config) (CrawlerInterface, error) {
	switch crawlerType {
	case DivarCrawlerType:
		return &DivarCrawler{config: config}, nil
	case SheypoorCrawlerType:
		return &SheypoorCrawler{config: config}, nil
	default:
		return nil, errors.New("invalid crawler type")
	}
}
