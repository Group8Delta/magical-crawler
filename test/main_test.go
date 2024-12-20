package test

import (
	// "magical-crwler/config"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/crawler"
	"magical-crwler/services/watchList"
	"os"
	"testing"
)

var testDbService database.DbService
var testRepo database.IRepository

var testDivarCrawler crawler.CrawlerInterface
var testSheypoorCrawler crawler.CrawlerInterface

var testWatchListService *watchList.WatchList

func TestMain(m *testing.M) {
	testDbService = database.New()
	config := config.GetConfig()
	err := testDbService.Init(config)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}
	testRepo = database.NewRepository(testDbService)

	defer testDbService.Close()

	divarCrawler, err := crawler.New(crawler.DivarCrawlerType, config, testRepo, 1, nil)
	if err != nil {
		panic("Failed to initial divar Crawler: " + err.Error())
	}
	testDivarCrawler = divarCrawler

	sheypoorCrawler, err := crawler.New(crawler.SheypoorCrawlerType, config, testRepo, 1, nil)
	if err != nil {
		panic("Failed to initial sheypoor Crawler: " + err.Error())
	}
	testSheypoorCrawler = sheypoorCrawler

	testWatchListService = watchList.New(testRepo, nil)

	code := m.Run()

	os.Exit(code)

}
