package main

import (
	"fmt"
	"log"
	"magical-crwler/config"
	"magical-crwler/database"
	"magical-crwler/services/FilterServices"
	"magical-crwler/services/alerting"
	"magical-crwler/services/bot"
	"magical-crwler/services/crawler"
	"magical-crwler/services/watchList"
	"os"
	"time"
)

func main() {
	conf := config.GetConfig()

	dbService := database.New()
	dbService.Init(conf)
	defer dbService.Close()

	db, err := dbService.GetDb().DB()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}
	err = db.Ping()
	if err != nil {
		fmt.Println("database connection error", err)
		os.Exit(1)

	}
	repo := database.NewRepository(dbService)

	err = setAdminTelegramIds(repo)
	if err != nil {
		fmt.Println("set admins had error:", err)
		os.Exit(1)
	}
	// I commented on this part because it needs a VPN to run
	bot, err := bot.NewBot(repo, bot.BotConfig{
		Token:  conf.BotToken,
		Poller: 10 * time.Second,
	})
	if err != nil {
		log.Println(err.Error())
	}

	alerter := alerting.NewAlerter(conf, bot)
	alerter.RunAdminNotifier()

	initialCrawlers(conf, repo, alerter)
	filterService := FilterServices.NewFilterServices(repo, bot)
	runFilterRunner(*filterService)

	watchListService := watchList.New(repo, bot)
	go watchListService.RunWatcher()

	bot.StartBot(dbService.GetDb())
	// http.ListenAndServe(":"+config.Port, nil)
}

func initialCrawlers(config *config.Config, repo database.IRepository, alerter *alerting.Alerter) {
	runIncrementalCrawl(config, repo, alerter)
	if config.EnableFullCrawl {
		runCrawlers(config, repo, 0, alerter)
		fmt.Println("full crawl started")
	}
}

func runCrawlers(c *config.Config, repo database.IRepository, maxDeepth int, alerter *alerting.Alerter) {
	for _, v := range crawler.CrawlerTypes {
		crawler, err := crawler.New(v, c, repo, maxDeepth, alerter)
		if err != nil {
			panic("Failed to initial Crawler: " + err.Error())
		}
		go crawler.RunCrawler()

	}
}
func runIncrementalCrawl(c *config.Config, repo database.IRepository, alerter *alerting.Alerter) {
	go func() {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				err := setAdminTelegramIds(repo)
				if err != nil {
					fmt.Println("set admins had error:", err)
				}
				runCrawlers(c, repo, 1, alerter)

			}
		}
	}()
}

func runFilterRunner(filterService FilterServices.FilterServices) {
	go func(filterService FilterServices.FilterServices) {
		ticker := time.NewTicker(2 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				filterService.ApplyFilters()
			}
		}
	}(filterService)
}

func setAdminTelegramIds(repo database.IRepository) error {
	users, err := repo.GetAdminUsers()
	if err != nil {
		return err
	}
	config.AdminTelegramIds = []int{}
	for _, v := range users {
		config.AdminTelegramIds = append(config.AdminTelegramIds, int(v.TelegramID))
	}
	return nil
}
