package crawler

import (
	"context"
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
	CrawlAdsLinks(ctx context.Context,searchUrl string) ([]string, error)
	CrawlPageUrl(ctx context.Context,pageUrl string) (*Ad, error)
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
