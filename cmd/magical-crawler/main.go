package main

import (
	"fmt"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/models"
	"magical-crwler/services/alerting"
	"magical-crwler/services/bot"
	"magical-crwler/services/crawler"
	"os"
	"time"
)

func main() {
	conf := config.GetConfig()

	database := database.New()
	database.Init(conf)
	defer database.Close()

	db, err := database.GetDb().DB()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}

	err = setAdminUserIds(database)
	fmt.Println(config.AdminUserIds)
	if err != nil {
		fmt.Println("set admins had error:", err)
		os.Exit(1)
	}
	// I commented on this part because it needs a VPN to run
	bot, err := bot.NewBot(bot.BotConfig{
		Token:  conf.BotToken,
		Poller: 10 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}

	alerter := alerting.NewAlerter(conf, bot)
	alerter.RunAdminNotifier()

	initialCrawlers(conf, database, alerter)

	bot.StartBot()
	// http.ListenAndServe(":"+config.Port, nil)
}

func initialCrawlers(config *config.Config, database database.DbService, alerter *alerting.Alerter) {
	runIncrementalCrawl(config, database, alerter)
	if config.EnableFullCrawl {
		runCrawlers(config, 0, alerter)
		fmt.Println("full crawl started")
	}
}

func runCrawlers(c *config.Config, maxDeepth int, alerter *alerting.Alerter) {
	for _, v := range crawler.CrawlerTypes {
		crawler, err := crawler.New(v, c, maxDeepth, alerter)
		if err != nil {
			panic("Failed to initial Crawler: " + err.Error())
		}
		crawler.RunCrawler()

	}
}
func runIncrementalCrawl(c *config.Config, database database.DbService, alerter *alerting.Alerter) {
	go func() {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := setAdminUserIds(database)
				if err != nil {
					fmt.Println("set admins had error:", err)
				}
				runCrawlers(c, 1, alerter)

			}
		}
	}()
}

func setAdminUserIds(database database.DbService) error {
	gormDb := database.GetDb()

	var users []*models.User
	result := gormDb.Where("role_id < ?", "3").Find(&users)
	if result.Error != nil {
		return result.Error
	}
	config.AdminUserIds = []int{}
	for _, v := range users {
		config.AdminUserIds = append(config.AdminUserIds, int(v.ID))
	}
	return nil
}
