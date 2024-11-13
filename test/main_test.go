package test

import (
	// "magical-crwler/config"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/crawler"
	"os"
	"testing"
)

var testDbService database.DbService
var testRepo *database.Repository

var testDivarCrawler crawler.CrawlerInterface
var testSheypoorCrawler crawler.CrawlerInterface

func TestMain(m *testing.M) {
	testDbService = database.New()
	config := config.GetConfig()
	testRepo := database.NewRepository(testDbService)
	err := testDbService.Init(config)
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

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

	code := m.Run()

	os.Exit(code)

}
