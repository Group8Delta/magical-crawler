package models

type CrawlerSetting struct {
	ID                               uint
	CrawlTimeOutPerSearchUrlInSecond int
	PageNumberLimit                  int
}
