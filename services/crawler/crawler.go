package crawler

import "errors"

type CrawlerType string

const (
	DivarCrawlerType    CrawlerType = "divar_crawler"
	SheypoorCrawlerType CrawlerType = "sheypoor_crawler"
)

type CrawlerInterface interface {
	CrawlAdsLinks(searchUrl string) ([]string, error)
	CrawlPageUrl(pageUrl string) (interface{}, error)
}

func New(crawlerType CrawlerType) (CrawlerInterface, error) {
	switch crawlerType {
	case DivarCrawlerType:
		return &DivarCrawler{}, nil
	case SheypoorCrawlerType:
		return &SheypoorCrawler{}, nil
	default:
		return nil, errors.New("invalid crawler type")
	}
}
