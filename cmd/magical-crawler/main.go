package main

import (
	"fmt"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/bot"
	"magical-crwler/services/crawler"
	"time"
)

func init() {
	config := config.GetConfig()
	runIncrementalCrawl(config)
	if config.EnableFullCrawl {
		runCrawlers(config, 0)
		fmt.Println("full crawl started")
	}
}

func runCrawlers(c *config.Config, maxDeepth int) {
	for _, v := range crawler.CrawlerTypes {
		crawler, err := crawler.New(v, c, maxDeepth)
		if err != nil {
			panic("Failed to initial Crawler: " + err.Error())
		}
		crawler.RunCrawler()

	}
}
func runIncrementalCrawl(c *config.Config) {
	go func() {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				runCrawlers(c, 1)

			}
		}
	}()
}
func main() {
	config := config.GetConfig()

	database := database.New()
	database.Init(config)
	defer database.Close()

	gdb := database.GetDb()
	db, err := gdb.DB()
	if err != nil {
		fmt.Println("database connection error", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database connection error", err)
	}
	// I commented on this part because it needs a VPN to run
	bot, err := bot.NewBot(bot.BotConfig{
		Token:  config.BotToken,
		Poller: 10 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}
	bot.StartBot(gdb)
	// http.ListenAndServe(":"+config.Port, nil)
}
